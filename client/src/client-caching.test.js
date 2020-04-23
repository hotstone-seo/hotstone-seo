import { HotStoneClient } from './index'
import nock from 'nock'
import crypto from 'crypto'

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
        });
        mockServer = nock(baseURL)
    })

    afterEach(() => {
        console.log('pending mocks: %j', mockServer.pendingMocks())
        mockServer.done()
    })

    describe('tags', () => {
        test('uses Expires header if no Pragma or Cache-Control', async () => {
            const mockResp = [
                { id: 1, type: "title" },
                { id: 2, type: "meta" },
            ]
            const givenRule = { rule_id: 9, path_param: { src: 'JKTC', dst: 'MESC' } }
            const givenLocale = 'en_US'
            
            mockServer
                .get('/p/fetch-tags')
                .query({ _rule: 9, _locale: 'en_US', src: 'JKTC', dst: 'MESC' })
                .reply(200, mockResp, {
                    // 'Cache-Control': 'max-age=30',
                    // 'Cache-Control': 'no-cache',
                    'Expires': new Date(new Date() - 1000).toUTCString(),
                    'Date': new Date().toUTCString(),
                    // 'ETag': 'deadbeef',
                    'Last-Modified': new Date(new Date() - 10000000).toUTCString(),

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
                    //TODO: assert If-Modified-Since
                    console.log('>>> REQ HEADERS: ', this.req.headers)
                })

            const tags2 = await subject.tags(givenRule, givenLocale)
            expect(tags2).toEqual(mockResp)
        })
    })
})