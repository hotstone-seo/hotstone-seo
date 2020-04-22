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

    describe('match', () => {
        test('good response', async () => {
            const mockResp = {
                rule_id: 1,
                path_param: {}
            }
        
            mockServer.get('/p/match')
            .query({_path: '/bar/fred'})
            .reply(200, mockResp)
    
            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual(mockResp)
        })
    
        test('bad response', async () => {
            mockServer.get('/p/match')
            .query({_path: '/bar/fred'})
            .reply(400)
    
            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual({})
        })
    })
})