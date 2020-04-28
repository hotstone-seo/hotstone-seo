import { HotStoneClient } from './index'
import nock from 'nock'
import crypto from 'crypto'
import RedisCache from './redis-cache'

function random(len) {
    return crypto
        .randomBytes(Math.ceil(len / 2))
        .toString('hex') // convert to hexadecimal format
        .slice(0, len)
}

describe('HotStone-Client with Caching', () => {
    let subject;
    let mockServer;

    beforeEach(() => {
        const baseURL = "http://foo.com"
        subject = new HotStoneClient(baseURL, {
            cacheManager: `./test-local-cache/${random(7)}`,
            // cacheManager: new RedisCache({ prefix: `test-${random(7)}:`, host: 'localhost', port: 6379, password: 'redispass' })
        });
        mockServer = nock(baseURL)
    })

    afterEach(() => {
        console.log('pending mocks: %j', mockServer.pendingMocks())

        // CAUTION: If you got failed test with message "Mocks not yet satisfied:" and pointing below line ('mockServer.done()'),
        // it means there is a failed assertion `expect` inside nock `reply` callback function.
        // Fix that. Don't ever comment below line to temporarily fix the issue.
        mockServer.done()
    })

    // see test case example: https://github.com/npm/make-fetch-happen/blob/latest/test/cache.js#L895
    describe('tags', () => {
        test('uses Expires header if no Pragma or Cache-Control', async () => {
            const mockResp = [
                { id: 1, type: "title" },
                { id: 2, type: "meta" },
            ]
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'

            const date = new Date()
            const expires = new Date(date - 1000)
            const lastModified = new Date(date - 10000000)

            mockServer
                .get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(200, mockResp, {
                    // 'Cache-Control': 'max-age=30',
                    // 'Cache-Control': 'no-cache',
                    'Expires': expires.toUTCString(),
                    'Date': date.toUTCString(),
                    // 'ETag': 'deadbeef',
                    'Last-Modified': lastModified.toUTCString(),
                    // 'Cache-Control': 'max-age=30',
                    // 'Expires': 'Thu, 23 Apr 2020 09:33:40 GMT',
                    // 'Last-Modified': 'Thu, 23 Apr 2020 09:33:10 GMT',
                    // 'Date': 'Thu, 23 Apr 2020 09:33:12 GMT',
                })

            const tags1 = await subject.tags(givenRule, givenLocale)
            expect(tags1).toEqual(mockResp)

            mockServer
                .get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(304, function () {
                    expect(this.req.headers['if-modified-since'][0]).toEqual(lastModified.toUTCString())
                })

            const tags2 = await subject.tags(givenRule, givenLocale)
            expect(tags2).toEqual(mockResp)
        })

        test('GIVEN Cache-Control max-age header THEN use local cache for 2nd resp without contact upstream server', async () => {
            const mockResp = [
                { id: 1, type: "title" },
                { id: 2, type: "meta" },
            ]
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'

            const age = 30
            const date = new Date()
            const lastModified = new Date(date)
            const expires = new Date(lastModified + (age * 1000))

            mockServer
                .get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(200, mockResp, {
                    'Cache-Control': `max-age=${age}`,
                    'Expires': expires.toUTCString(),
                    'Date': date.toUTCString(),
                    'Last-Modified': lastModified.toUTCString(),
                })

            const tags1 = await subject.tags(givenRule, givenLocale)
            expect(tags1).toEqual(mockResp)

            const tags2 = await subject.tags(givenRule, givenLocale)
            expect(tags2).toEqual(mockResp)
        })
    })
})