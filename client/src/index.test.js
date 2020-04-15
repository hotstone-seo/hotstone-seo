import { HotStoneClient } from './index'
import { Server, Response } from "miragejs"

describe('HotStone-Client', () => {
    let subject;
    let mockServer;

    beforeEach(() => {
        subject = new HotStoneClient('https://foo.com');
        mockServer = new Server({
            urlPrefix: 'https://foo.com',
            trackRequests: true
        })
    })
    afterEach(() => {
        mockServer.shutdown();
    });

    describe('match', () => {
        test('good response', async () => {
            const mockResp = {
                rule_id: 1,
                path_param: {}
            }
            mockServer.post("/p/match", (schema, request) => { return mockResp })

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual(mockResp)

            const requests = mockServer.pretender.handledRequests
            expect(requests.length).toBe(1)
            expect(requests[0].requestBody).toBe(JSON.stringify({ path: '/bar/fred' }))
        })

        test('bad response', async () => {
            mockServer.post("/p/match", (schema, request) => { 
                return new Response(400);
            })

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual({})

            const requests = mockServer.pretender.handledRequests
            expect(requests.length).toBe(1)
            expect(requests[0].requestBody).toBe(JSON.stringify({ path: '/bar/fred' }))
        })
    })
})
