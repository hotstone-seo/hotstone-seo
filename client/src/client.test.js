import { HotStoneClient } from './index'
import nock from 'nock'

describe('HotStone-Client', () => {
    let subject;
    let mockServer;

    beforeEach(() => {
        const baseURL = "http://foo.com"
        subject = new HotStoneClient(baseURL);
        mockServer = nock(baseURL)
    })

    afterEach(() => {
        mockServer.done()
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

    })
})