import { HotStoneClient } from './index'
import axios from 'axios';
import MockAdapter from 'axios-mock-adapter'

describe('HotStone-Client', () => {
    let mockApiCaller;
    let subject;

    beforeEach(() => {
        mockApiCaller = new MockAdapter(axios);
        subject = new HotStoneClient('https://foo.com');
    })
    afterEach(() => {
        mockApiCaller.restore();
    });

    describe('match', () => {
        test('good response', async () => {
            const mockResp = {
                rule_id: 1,
                path_param: {}
            }

            mockApiCaller.onPost('/p/match', { path: '/bar/fred' }).replyOnce(200, mockResp)

            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual(mockResp)

            expect(mockApiCaller.history.post.length).toBe(1);
            expect(mockApiCaller.history.post[0].data).toBe(JSON.stringify({ path: '/bar/fred' }));
        })

        test('bad response', async () => {
            mockApiCaller.onPost('/p/match', { path: '/bar/fred' }).replyOnce(404)
            const rule = await subject.match('/bar/fred')
            expect(rule).toEqual({})
        })
    })
})
