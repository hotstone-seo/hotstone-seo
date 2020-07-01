import { HotStoneClient } from './index'
import nock from 'nock'

describe('HotStone-Client', () => {
    let subject;
    let mockServer;

    beforeEach(() => {
        const baseURL = "http://foo.com"
        const key = "abc.12345678901234567890123456789012"
        subject = new HotStoneClient(baseURL, key);
        mockServer = nock(baseURL)
    })

    afterEach(() => {
        mockServer.done()
    })

    describe('baseURL', () => {
        const testCases = [
            { name: 'no trailing slash', givenBaseURL: 'http://foo.com', expectBaseURL: 'http://foo.com' },
            { name: 'single trailing slash', givenBaseURL: 'http://foo.com/', expectBaseURL: 'http://foo.com' },
            { name: 'multi trailing slash', givenBaseURL: 'http://foo.com///', expectBaseURL: 'http://foo.com' },
            { name: 'only trailing slash', givenBaseURL: '/', expectBaseURL: '/' }
        ]

        testCases.map(({ name, givenBaseURL, expectBaseURL }) => {
            test(name, async () => {
                const key = "abc.12345678901234567890123456789012"
                const client = new HotStoneClient(givenBaseURL, key);
                expect(client.baseURL).toEqual(expectBaseURL)
            })
        })
    })

    describe('match', () => {
        test('good response', async () => {
            const mockResp = {
                rule_id: 1,
                path_param: {}
            }

            mockServer.get('/p/match')
                .query({ _path: '/bar/fred' })
                .reply(200, mockResp)

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual(mockResp)
        })

        test('bad response', async () => {
            mockServer.get('/p/match')
                .query({ _path: '/bar/fred' })
                .reply(400)

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual({})
        })

        test('bad format response', async () => {
            const mockResp = {
                message: 'foo msg'
            }

            mockServer.get('/p/match')
                .query({ _path: '/bar/fred' })
                .reply(200, mockResp)

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual({})
        })

        test('not json format', async () => {
            mockServer.get('/p/match')
                .query({ _path: '/bar/fred' })
                .reply(200, `<html>message</html>`)

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual({})
        })
    })

    describe('tags', () => {
        test('good response', async () => {
            const mockResp = [
                { id: 1, type: "title" },
                { id: 2, type: "meta" },
            ]
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'

            mockServer.get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(200, mockResp)

            const tags = await subject.tags(givenRule, givenLocale)
            expect(tags).toEqual(mockResp)
        })

        test('bad response', async () => {
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'

            mockServer.get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(400)

            const tags = await subject.tags(givenRule, givenLocale)
            expect(tags).toEqual([])
        })

        test('bad format response', async () => {
            const mockResp = {
                message: 'foo msg'
            }
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'

            mockServer.get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(200, mockResp)

            const tags = await subject.tags(givenRule, givenLocale)
            expect(tags).toEqual(mockResp)
        })

        test('not json format', async () => {
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'

            mockServer.get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(200, `<html>message</html>`)

            const tags = await subject.tags(givenRule, givenLocale)
            expect(tags).toEqual([])
        })

    })
})