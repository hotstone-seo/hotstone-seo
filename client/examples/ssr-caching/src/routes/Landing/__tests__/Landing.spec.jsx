import React from 'react';
import sinon from 'sinon';
import Landing from '../';
import Breadcrumbs from '../../../components/Breadcrumbs';
import OrderDetailPanel from '../../../components/OrderDetailPanel';
import PaymentSelector from '../components/PaymentSelector';
import * as payment from '../../../models/payment';

const props = {
  paymentList: {
    paymentDescription: {
      currency: 'IDR',
      orderId: '39981775',
      orderHash: 'ac526ba985927bb0b846a79c9d0be36321db721b',
      grandTotal: 164996,
      grandSubTotal: 165000,
      paymentCharge: 0,
      uniqueCode: -4,
      orderExpiredDatetime: '2018-03-28 18:48:12',
      paymentType: 3,
      tixPoint: 0,
      baggageFee: [100, 200],
      giftPromo: [
        {
          order_detail_id: '1',
          order_name: 'test',
          selling_price: 3000,
          order_type: 'flight'
        }
      ]
    }
  },
  account: {
    point_balance: {
      point_amount: 100
    }
  },
  query: {
    order_id: '39982221',
    order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'
  }
};

test('Should render Landing correctly', () => {
  const tree = enzyme.shallow(<Landing paymentList={props.paymentList} query={props.query} />);

  expect(tree).toMatchSnapshot();
});

test('Should contain Breadcrumbs, OrderDetailPanel, PaymentSelector', () => {
  const wrapper = mount(<Landing paymentList={props.paymentList} query={props.query} />);

  expect(wrapper.find(Breadcrumbs)).toHaveLength(1);
  expect(wrapper.find(OrderDetailPanel)).toHaveLength(1);
  expect(wrapper.find(PaymentSelector)).toHaveLength(1);
});

test('Should display expire page error correctly', () => {
  const availablePaymentAction = jest.fn();
  const mounting = () => {
    mount(
      <Landing.WrappedComponent
        paymentList={{
          ...props.paymentList,
          error_msgs: 'kedaluwarsa test'
        }}
        getAvailablePaymentAction={availablePaymentAction}
        account={props.account}
        query={props.query}
      />
    );
  };

  expect(mounting).toThrowError('Pesanan ini sudah Kadaluarsa, Harap lakukan pemesanan baru');
});

test('Should display error page correctly', () => {
  const availablePaymentAction = jest.fn();
  const mounting = () => {
    mount(
      <Landing.WrappedComponent
        paymentList={{
          ...props.paymentList,
          error_msgs: 'test'
        }}
        getAvailablePaymentAction={availablePaymentAction}
        account={props.account}
        query={props.query}
      />
    );
  };

  expect(mounting).toThrowError('test');
});

test('Should handle all tix point functions', () => {
  //noinspection JSAnnotator
  payment.getUseTix = jest.fn();
  //noinspection JSAnnotator
  payment.getCancelPromo = jest.fn();
  payment.getCancelPromo
    .mockReturnValueOnce(Promise.resolve({ data: { diagnostic: { error_msgs: 'test' } } }))
    .mockReturnValueOnce(Promise.resolve({ data: { diagnostic: { confirm: 'success' } } }));
  payment.getUseTix
    .mockReturnValueOnce(Promise.resolve({ data: { diagnostic: { error_msgs: 'test' } } }))
    .mockReturnValueOnce(Promise.resolve({ data: { diagnostic: { confirm: 'success' } } }));
  const availablePaymentAction = jest.fn();
  const resetPopup = jest.fn();
  const setPopup = jest.fn();

  sinon.spy(Landing.WrappedComponent.prototype, 'handleShowTIXDialog');
  sinon.spy(Landing.WrappedComponent.prototype, 'handleHideTIXDialog');
  sinon.spy(Landing.WrappedComponent.prototype, 'handleUseTIX');
  sinon.spy(Landing.WrappedComponent.prototype, 'handleCancelPromo');
  let wrapper = mount(
    <Landing.WrappedComponent
      paymentList={{
        ...props.paymentList,
        paymentDescription: { ...props.paymentDescription, giftPromo: false }
      }}
      account={props.account}
      getAvailablePaymentAction={availablePaymentAction}
      resetPopup={resetPopup}
      setPopup={setPopup}
      query={props.query}
    />
  );

  expect(availablePaymentAction.mock.calls.length).toEqual(1);
  console.log(wrapper.debug());
  // Test use tix display dialog
  wrapper.find('.payment-tixpoint-usage').simulate('click');
  expect(Landing.WrappedComponent.prototype.handleShowTIXDialog.calledOnce).toEqual(true);

  // Test use tix cancel
  wrapper.find('a.hide-tix-modal').simulate('click');
  expect(Landing.WrappedComponent.prototype.handleHideTIXDialog.calledOnce).toEqual(true);

  // Test use tix fail
  wrapper.find('.payment-tixpoint-usage').simulate('click');
  wrapper.find('a.use-tix').simulate('click');
  expect(Landing.WrappedComponent.prototype.handleUseTIX.calledOnce).toEqual(true);

  // Test use tix success
  wrapper.find('.payment-tixpoint-usage').simulate('click');
  wrapper.find('a.use-tix').simulate('click');

  wrapper.unmount();
  wrapper = mount(
    <Landing.WrappedComponent
      paymentList={props.paymentList}
      account={props.account}
      getAvailablePaymentAction={availablePaymentAction}
      resetPopup={resetPopup}
      setPopup={setPopup}
      query={props.query}
    />
  );
  // Test additional payment cancel fail
  wrapper.find('.additional-payment-cancel').simulate('click');
  expect(Landing.WrappedComponent.prototype.handleCancelPromo.calledOnce).toEqual(true);

  // Test additional payment cancel success
  wrapper.find('.additional-payment-cancel').simulate('click');

  expect(availablePaymentAction.mock.calls.length).toEqual(2);
});
