import reducers from '../reducers';
import * as account from '../../models/account';
import * as payment from '../../models/payment';

test('reducers', () => {
  //noinspection JSAnnotator
  account.getAccount = jest.fn();
  account.getAccount.mockReturnValue(Promise.resolve({data: {}}));

  //noinspection JSAnnotator
  payment.getAvailablePayment = jest.fn();
  payment.getAvailablePayment.mockReturnValue(Promise.resolve({data: {}}));

  //noinspection JSAnnotator
  payment.getPaymentConfirmation = jest.fn();
  payment.getPaymentConfirmation.mockReturnValue(Promise.resolve({data: {}}));

  //noinspection JSAnnotator
  payment.getConfirmPaymentIndex = jest.fn();
  payment.getConfirmPaymentIndex.mockReturnValue(Promise.resolve({data: {}}));

  //noinspection JSAnnotator
  payment.getPaymentDetail = jest.fn();
  payment.getPaymentDetail.mockReturnValue(Promise.resolve({data: {}}));

  //noinspection JSAnnotator
  payment.getOrderDetail = jest.fn();
  payment.getOrderDetail.mockReturnValue(Promise.resolve({data: {}}));

  let state;

  state = reducers(undefined, {});
  console.log(JSON.stringify(state));
  expect(state).toEqual({
      payment: {
        groupDetail: {loading: false, loaded: false,result: {}},
        paymentList: {
          paymentMethods: [],
          paymentDescription: {},
          loading: false,
          loaded: false
        },
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: false, loaded: false, data: {}},
        context: {query: {}, params: {}, isWebView: false}
      }
    }
  );
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: false, loaded: false, data: {}},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_PAYMENT_DETAIL'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: true, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {loading: false, loaded: false, data: {}},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: false, loaded: false, data: {}},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_ACCOUNT'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: true, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {loading: true, loaded: false, data: {}},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: true, loaded: false, data: {}},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {result: {}, type: 'LOAD_ACCOUNT_SUCCESS'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: true, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: true, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: true, loaded: false, data: {}},
        context: {
          query: {order_id: '39982384', order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc'},
          params: {}
        }
      }
    },
    {
      error: {
        config: {
          transformRequest: {},
          transformResponse: {},
          timeout: 0,
          xsrfCookieName: 'XSRF-TOKEN',
          xsrfHeaderName: 'X-XSRF-TOKEN',
          maxContentLength: 65536,
          headers: {
            Accept: 'application/json, text/plain, */*',
            'Content-Type': 'application/x-www-form-urlencoded',
            TIXAPI: 1
          },
          method: 'get',
          url: 'https://renan.tiket.com/myaccount/mypoints',
          data: '',
          maxRedirects: 0,
          responseType: 'json',
          withCredentials: true
        },
        request: {}
      },
      type: 'LOAD_ACCOUNT_FAIL'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: true, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {
        data: {},
        error: {
          config: {
            transformRequest: {},
            transformResponse: {},
            timeout: 0,
            xsrfCookieName: 'XSRF-TOKEN',
            xsrfHeaderName: 'X-XSRF-TOKEN',
            maxContentLength: 65536,
            headers: {
              Accept: 'application/json, text/plain, */*',
              'Content-Type': 'application/x-www-form-urlencoded',
              TIXAPI: 1
            },
            method: 'get',
            url: 'https://renan.tiket.com/myaccount/mypoints',
            data: '',
            maxRedirects: 0,
            responseType: 'json',
            withCredentials: true
          },
          request: {}
        },
        loading: false,
        loaded: false
      },
      context: {
        query: {order_id: '39982384', order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {
      result: {
        output_type: 'json',
        diagnostic: {
          status: 200,
          elapsetime: '0.8285',
          memoryusage: '41.63MB',
          unix_timestamp: 1522815480,
          lang: 'id',
          currency: 'IDR'
        },
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 2,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_transfer_multiple_payment: [
            {
              code: '12',
              type: 'bca_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'BCA Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              num_type: 12,
              error_message: '',
              icon: 'payment-transfer-bca.png',
              is_disabled: '',
              class: 'payment-bca-5',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            },
            {
              code: '42',
              type: 'mandiri_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Mandiri Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              num_type: 42,
              error_message: '',
              icon: 'payment-transfer-mandiri.png',
              is_disabled: '',
              class: 'payment-mandiri-1',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            }
          ]
        },
        login_status: 'true',
        guest_id: '20661143',
        login_email: 'bayu@tiket.com',
        token: '089bcde8b90b8ff4343d6f99125b480b84cbc6c7'
      },
      type: 'LOAD_PAYMENT_DETAIL_SUCCESS'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 2,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_transfer_multiple_payment: [
            {
              code: '12',
              type: 'bca_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'BCA Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              num_type: 12,
              error_message: '',
              icon: 'payment-transfer-bca.png',
              is_disabled: '',
              class: 'payment-bca-5',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            },
            {
              code: '42',
              type: 'mandiri_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Mandiri Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              num_type: 42,
              error_message: '',
              icon: 'payment-transfer-mandiri.png',
              is_disabled: '',
              class: 'payment-mandiri-1',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            }
          ],
          token: '089bcde8b90b8ff4343d6f99125b480b84cbc6c7'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {
      result: {
        output_type: 'json',
        diagnostic: {
          status: 200,
          elapsetime: '0.8285',
          memoryusage: '41.63MB',
          unix_timestamp: 1522815480,
          lang: 'id',
          error_msgs: 'test',
          currency: 'IDR'
        },
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 2,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_transfer_multiple_payment: [
            {
              code: '12',
              type: 'bca_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'BCA Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              num_type: 12,
              error_message: '',
              icon: 'payment-transfer-bca.png',
              is_disabled: '',
              class: 'payment-bca-5',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            },
            {
              code: '42',
              type: 'mandiri_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Mandiri Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              num_type: 42,
              error_message: '',
              icon: 'payment-transfer-mandiri.png',
              is_disabled: '',
              class: 'payment-mandiri-1',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            }
          ]
        },
        login_status: 'true',
        guest_id: '20661143',
        login_email: 'bayu@tiket.com',
        token: '089bcde8b90b8ff4343d6f99125b480b84cbc6c7'
      },
      type: 'LOAD_PAYMENT_DETAIL_SUCCESS'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 2,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_transfer_multiple_payment: [
            {
              code: '12',
              type: 'bca_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'BCA Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              num_type: 12,
              error_message: '',
              icon: 'payment-transfer-bca.png',
              is_disabled: '',
              class: 'payment-bca-5',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            },
            {
              code: '42',
              type: 'mandiri_transfer',
              link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Mandiri Transfer',
              id: '',
              message: '',
              images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              num_type: 42,
              error_message: '',
              icon: 'payment-transfer-mandiri.png',
              is_disabled: '',
              class: 'payment-mandiri-1',
              desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
              tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
            }
          ],
          token: '089bcde8b90b8ff4343d6f99125b480b84cbc6c7',
          error_msgs: 'test'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: false, loaded: false, data: {}},
        context: {
          query: {order_id: '39982384', order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc'},
          params: {}
        }
      }
    },
    {type: 'LOAD_AVAILABLE_PAYMENTS'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: true, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {loading: false, loaded: false, data: {}},
      context: {
        query: {order_id: '39982384', order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: true, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {
          data: {},
          error: {
            config: {
              transformRequest: {},
              transformResponse: {},
              timeout: 0,
              xsrfCookieName: 'XSRF-TOKEN',
              xsrfHeaderName: 'X-XSRF-TOKEN',
              maxContentLength: 65536,
              headers: {
                Accept: 'application/json, text/plain, */*',
                'Content-Type': 'application/x-www-form-urlencoded',
                TIXAPI: 1
              },
              method: 'get',
              url: 'https://renan.tiket.com/myaccount/mypoints',
              data: '',
              maxRedirects: 0,
              responseType: 'json',
              withCredentials: true
            },
            request: {}
          },
          loading: false,
          loaded: false
        },
        context: {
          query: {order_id: '39982384', order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc'},
          params: {}
        }
      }
    },
    {
      result: {
        output_type: 'json',
        diagnostic: {
          status: 200,
          elapsetime: '0.7937',
          memoryusage: '41.64MB',
          unix_timestamp: 1523261020,
          confirm: 'success',
          lang: 'id',
          currency: 'IDR'
        },
        sidebar_view: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982384',
          order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7014,
          grand_total: 10014,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982384',
              order_detail_id: '58944666',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-09 18:35:42',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-09 14:35:42',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '1347070642693',
              checkin_date: null,
              tiket_cust_name: 'Ihsan Fauzi Rahman',
              tiket_gender: 'f',
              tiket_no_hp: '+6281911776566',
              tiket_birth_date: '1996-02-21',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: 'Ihsan Fauzi Rahman',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '24618237',
              tiket_detail: [
                {
                  tiket_barcode: '1347070642693',
                  tiket_cust_name: 'Ihsan Fauzi Rahman',
                  tiket_cust_id: 'Ihsan Fauzi Rahman',
                  tiket_seating: '',
                  tiket_gender: 'f'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Ihsan Fauzi Rahman',
                  id_card: 'Ihsan Fauzi Rahman',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 3,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-09 18:20:42',
          calculate_tixpoint_earn: {
            total_earn: 0,
            total_amount: 0,
            earn_ratio: 20,
            point_currency: 'IDR',
            point_breakdown: []
          },
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: ''
        },
        countdown_expired_datetime: 11822,
        available_payment: [
          {
            code: '1',
            link: 'https://renan.tiket.com/checkout/checkout_payment/1?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Kartu Kredit',
            message: '',
            type: 'creditcard',
            desc: 'Master Card and Visa',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/visa.png',
              'https://renan.tiket.com/images/apps_payment/master_card.png'
            ],
            countdown_time: 11822
          },
          {
            code: '2',
            link: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Transfer',
            message: '',
            type: 'banktransfer',
            desc: 'Transfer',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png'
            ],
            payment_group: [
              {
                code: '12',
                type: 'bca_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'BCA Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              },
              {
                code: '42',
                type: 'mandiri_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'Mandiri Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              }
            ],
            countdown_time: 11822
          },
          {
            code: '99',
            type: 'virtualaccount',
            link: 'https://renan.tiket.com/checkout/checkout_payment/99?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Virtual Account',
            message: '',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bri_transfer.png'
            ],
            desc: 'Virtual Account',
            payment_group: [
              {
                code: '13',
                type: 'va_bca',
                link: 'http://renan.tiket.com/checkout/checkout_payment/13?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA BCA',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              },
              {
                code: '43',
                type: 'va_mandiri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/43?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA Mandiri',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              },
              {
                code: '52',
                type: 'va_bni',
                link: 'http://renan.tiket.com/checkout/checkout_payment/52?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA BNI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking, Mobile Banking & SMS Banking',
                countdown_time: 11822
              },
              {
                code: '62',
                type: 'va_bri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/62?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA BRI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              }
            ],
            countdown_time: 11822
          },
          {
            code: '59',
            link: 'https://renan.tiket.com/checkout/checkout_payment/59?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'ATM',
            message: '',
            type: 'atm_nicepay',
            desc: 'ATM Bersama, Prima & Alto',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/atm_bersama.png',
              'https://renan.tiket.com/images/apps_payment/atm_prima.png',
              'https://renan.tiket.com/images/apps_payment/alto.png'
            ],
            countdown_time: 11822
          },
          {
            code: '3',
            link: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'KlikBCA',
            message: '',
            type: 'klikbca',
            desc: 'KlikBCA',
            app_images: ['https://renan.tiket.com/images/apps_payment/klik_bca.png'],
            countdown_time: 11822
          },
          {
            code: '4',
            link: 'https://renan.tiket.com/checkout/checkout_payment/4?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'BCA KlikPay',
            message: '',
            type: 'klikpay',
            desc: 'BCA Klikpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/bca_klikpay.png'],
            countdown_time: 11822
          },
          {
            code: '34',
            link: 'https://renan.tiket.com/checkout/checkout_payment/34?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Mandiri Clickpay',
            message: 'Butuh No. Kartu dan Token',
            type: 'mandiri_clickpay',
            desc: 'Mandiri Clickpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/mandiri_clickpay.png'],
            countdown_time: 11822
          },
          {
            code: '31',
            link: 'https://renan.tiket.com/checkout/checkout_payment/31?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'CIMB Clicks',
            message: '',
            type: 'cimbclicks',
            desc: 'CIMB Clicks',
            app_images: ['https://renan.tiket.com/images/apps_payment/cimb_click.png'],
            countdown_time: 11822
          },
          {
            code: '33',
            link: 'https://renan.tiket.com/checkout/checkout_payment/33?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'ePay BRI',
            message: '',
            type: 'epaybri',
            desc: 'ePay BRI',
            app_images: ['https://renan.tiket.com/images/apps_payment/e_pay_bri.png'],
            countdown_time: 11822
          }
        ],
        login_status: 'true',
        guest_id: '21699926',
        login_email: 'testing@tiket.com',
        token: 'e38c046d051cd2c6174207cc5abb11831ed8506f'
      },
      type: 'LOAD_AVAILABLE_PAYMENTS_SUCCESS'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {
        paymentMethods: [
          {
            code: '1',
            link: 'https://renan.tiket.com/checkout/checkout_payment/1?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Kartu Kredit',
            message: '',
            type: 'creditcard',
            desc: 'Master Card and Visa',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/visa.png',
              'https://renan.tiket.com/images/apps_payment/master_card.png'
            ],
            countdown_time: 11822
          },
          {
            code: '2',
            link: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Transfer',
            message: '',
            type: 'banktransfer',
            desc: 'Transfer',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png'
            ],
            payment_group: [
              {
                code: '12',
                type: 'bca_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'BCA Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              },
              {
                code: '42',
                type: 'mandiri_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'Mandiri Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              }
            ],
            countdown_time: 11822
          },
          {
            code: '99',
            type: 'virtualaccount',
            link: 'https://renan.tiket.com/checkout/checkout_payment/99?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Virtual Account',
            message: '',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bri_transfer.png'
            ],
            desc: 'Virtual Account',
            payment_group: [
              {
                code: '13',
                type: 'va_bca',
                link: 'http://renan.tiket.com/checkout/checkout_payment/13?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA BCA',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              },
              {
                code: '43',
                type: 'va_mandiri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/43?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA Mandiri',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              },
              {
                code: '52',
                type: 'va_bni',
                link: 'http://renan.tiket.com/checkout/checkout_payment/52?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA BNI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking, Mobile Banking & SMS Banking',
                countdown_time: 11822
              },
              {
                code: '62',
                type: 'va_bri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/62?order_id=39982384&order_hash8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
                text: 'VA BRI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 11822
              }
            ],
            countdown_time: 11822
          },
          {
            code: '59',
            link: 'https://renan.tiket.com/checkout/checkout_payment/59?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'ATM',
            message: '',
            type: 'atm_nicepay',
            desc: 'ATM Bersama, Prima & Alto',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/atm_bersama.png',
              'https://renan.tiket.com/images/apps_payment/atm_prima.png',
              'https://renan.tiket.com/images/apps_payment/alto.png'
            ],
            countdown_time: 11822
          },
          {
            code: '3',
            link: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'KlikBCA',
            message: '',
            type: 'klikbca',
            desc: 'KlikBCA',
            app_images: ['https://renan.tiket.com/images/apps_payment/klik_bca.png'],
            countdown_time: 11822
          },
          {
            code: '4',
            link: 'https://renan.tiket.com/checkout/checkout_payment/4?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'BCA KlikPay',
            message: '',
            type: 'klikpay',
            desc: 'BCA Klikpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/bca_klikpay.png'],
            countdown_time: 11822
          },
          {
            code: '34',
            link: 'https://renan.tiket.com/checkout/checkout_payment/34?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'Mandiri Clickpay',
            message: 'Butuh No. Kartu dan Token',
            type: 'mandiri_clickpay',
            desc: 'Mandiri Clickpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/mandiri_clickpay.png'],
            countdown_time: 11822
          },
          {
            code: '31',
            link: 'https://renan.tiket.com/checkout/checkout_payment/31?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'CIMB Clicks',
            message: '',
            type: 'cimbclicks',
            desc: 'CIMB Clicks',
            app_images: ['https://renan.tiket.com/images/apps_payment/cimb_click.png'],
            countdown_time: 11822
          },
          {
            code: '33',
            link: 'https://renan.tiket.com/checkout/checkout_payment/33?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
            text: 'ePay BRI',
            message: '',
            type: 'epaybri',
            desc: 'ePay BRI',
            app_images: ['https://renan.tiket.com/images/apps_payment/e_pay_bri.png'],
            countdown_time: 11822
          }
        ],
        paymentDescription: {
          currency: 'IDR',
          orderId: '39982384',
          orderHash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
          grandTotal: 10014,
          subTotal: 1000,
          grandSubTotal: 3000,
          paymentCharge: 2000,
          baggageFee: [0],
          uniqueCode: 7014,
          orderExpiredDatetime: '2018-04-09 18:20:42',
          paymentType: 3,
          tixPoint: 0,
          giftPromo: false
        },
        sidebar_view: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982384',
          order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7014,
          grand_total: 10014,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982384',
              order_detail_id: '58944666',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-09 18:35:42',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-09 14:35:42',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '1347070642693',
              checkin_date: null,
              tiket_cust_name: 'Ihsan Fauzi Rahman',
              tiket_gender: 'f',
              tiket_no_hp: '+6281911776566',
              tiket_birth_date: '1996-02-21',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: 'Ihsan Fauzi Rahman',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '24618237',
              tiket_detail: [
                {
                  tiket_barcode: '1347070642693',
                  tiket_cust_name: 'Ihsan Fauzi Rahman',
                  tiket_cust_id: 'Ihsan Fauzi Rahman',
                  tiket_seating: '',
                  tiket_gender: 'f'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Ihsan Fauzi Rahman',
                  id_card: 'Ihsan Fauzi Rahman',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 3,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982384&order_hash=8d1d895aca0f317dcb08ab840574d4fc2c95bfbc',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-09 18:20:42',
          calculate_tixpoint_earn: {
            total_earn: 0,
            total_amount: 0,
            earn_ratio: 20,
            point_currency: 'IDR',
            point_breakdown: []
          },
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: ''
        },
        loading: false,
        loaded: true
      },
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {
        data: {},
        error: {
          config: {
            transformRequest: {},
            transformResponse: {},
            timeout: 0,
            xsrfCookieName: 'XSRF-TOKEN',
            xsrfHeaderName: 'X-XSRF-TOKEN',
            maxContentLength: 65536,
            headers: {
              Accept: 'application/json, text/plain, */*',
              'Content-Type': 'application/x-www-form-urlencoded',
              TIXAPI: 1
            },
            method: 'get',
            url: 'https://renan.tiket.com/myaccount/mypoints',
            data: '',
            maxRedirects: 0,
            responseType: 'json',
            withCredentials: true
          },
          request: {}
        },
        loading: false,
        loaded: false
      },
      context: {
        query: {order_id: '39982384', order_hash: '8d1d895aca0f317dcb08ab840574d4fc2c95bfbc'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: '0',
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                total_price: 1000,
                quantity: '1',
                order_id: '39982221',
                order_detail_id: '58944361',
                order_detail_status: 'active',
                order_type: 'event',
                order_master_id: '4777',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                selling_price: '1000.00',
                order_expire_datetime: '2018-04-04 11:59:43',
                selling_currency: 'IDR',
                created_timestamp: '2018-04-04 09:46:47',
                order_name: 'Testing Event',
                order_name_detail: 'Test CC',
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                tiket_with_seating: '0',
                tiket_required_info: 'name,idcard',
                tiket_min_purchase: '1',
                tiket_max_purchase: '10',
                tiket_start_sell: '2018-03-01 00:00:00',
                tiket_end_sell: '2018-06-30 00:00:00',
                ext_source: 'native',
                ext_source_id: null,
                tiket_sell_price_netto_api: '0.00',
                tiket_id: '4777',
                tiket_total_allotment: '1000',
                tiket_markup_price_api: '0.00',
                tiket_subsidy_price_api: '0.00',
                uri: 'testing-event',
                file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
                business_id: '14161',
                business_address1: 'Lapangan D Senayan',
                country_name: 'Indonesia',
                city_name: 'Jakarta Selatan',
                voucher_provider: 'tiket.com',
                event_type: 'B',
                tiket_barcode: '6546216472445',
                checkin_date: null,
                tiket_cust_name: 'Abdul Rahman',
                tiket_gender: 'm',
                tiket_no_hp: '+6285692223310',
                tiket_birth_date: '1998-02-13',
                customer_price: '1000.00',
                customer_currency: 'IDR',
                tiket_cust_id: '09110110101',
                sell_rate_price: '1000.00',
                tiket_seating: '',
                tiket_attend_date: '0000-00-00',
                is_installment: '0',
                payment_status: 'shoppingcart',
                event_start: '2012-02-01 15:00:00',
                event_end: '2018-04-30 00:00:00',
                business_lat: '-6.20928550000000000000',
                business_long: '106.83307890000003000000',
                contact_person: '26042317',
                tiket_detail: [
                  {
                    tiket_barcode: '6546216472445',
                    tiket_cust_name: 'Abdul Rahman',
                    tiket_cust_id: '09110110101',
                    tiket_seating: '',
                    tiket_gender: 'm'
                  }
                ],
                tiket_quantity: 1,
                detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
                displayed_price: 1000,
                tax: 0,
                event_category: 'event',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                type: 'event',
                item_charge: 2000,
                item_charge_idr: 2000,
                tax_and_charge: 2000
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 2,
            is_confirmation: false,
            is_change_payment: false,
            type: false,
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            bank_transfer_multiple_payment: [
              {
                code: '12',
                type: 'bca_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'BCA Transfer',
                id: '',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                num_type: 12,
                error_message: '',
                icon: 'payment-transfer-bca.png',
                is_disabled: '',
                class: 'payment-bca-5',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
              },
              {
                code: '42',
                type: 'mandiri_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'Mandiri Transfer',
                id: '',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                num_type: 42,
                error_message: '',
                icon: 'payment-transfer-mandiri.png',
                is_disabled: '',
                class: 'payment-mandiri-1',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                tt_message: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                tt_message_m: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking'
              }
            ],
            token: '089bcde8b90b8ff4343d6f99125b480b84cbc6c7'
          },
          loading: false,
          loaded: true
        },
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'RESET_PAYMENT_DETAIL'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_ACCOUNT'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: true, loaded: false},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: true, loaded: false},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_PAYMENT_DETAIL'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: true, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: true, loaded: false},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: true, loaded: false},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {result: {}, type: 'LOAD_ACCOUNT_SUCCESS'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: true, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {
      result: {
        output_type: 'json',
        diagnostic: {
          status: 200,
          elapsetime: '0.8177',
          memoryusage: '41.52MB',
          unix_timestamp: 1522815482,
          lang: 'id',
          currency: 'IDR'
        },
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 12,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_image: '/images/ico_bca.png',
          important_information: [
            'Informasi Penting',
            'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
            'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
            'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
            'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
          ]
        },
        login_status: 'true',
        guest_id: '20661143',
        login_email: 'bayu@tiket.com',
        token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
      },
      type: 'LOAD_PAYMENT_DETAIL_SUCCESS'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 12,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_image: '/images/ico_bca.png',
          important_information: [
            'Informasi Penting',
            'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
            'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
            'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
            'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
          ],
          token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: '0',
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                total_price: 1000,
                quantity: '1',
                order_id: '39982221',
                order_detail_id: '58944361',
                order_detail_status: 'active',
                order_type: 'event',
                order_master_id: '4777',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                selling_price: '1000.00',
                order_expire_datetime: '2018-04-04 11:59:43',
                selling_currency: 'IDR',
                created_timestamp: '2018-04-04 09:46:47',
                order_name: 'Testing Event',
                order_name_detail: 'Test CC',
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                tiket_with_seating: '0',
                tiket_required_info: 'name,idcard',
                tiket_min_purchase: '1',
                tiket_max_purchase: '10',
                tiket_start_sell: '2018-03-01 00:00:00',
                tiket_end_sell: '2018-06-30 00:00:00',
                ext_source: 'native',
                ext_source_id: null,
                tiket_sell_price_netto_api: '0.00',
                tiket_id: '4777',
                tiket_total_allotment: '1000',
                tiket_markup_price_api: '0.00',
                tiket_subsidy_price_api: '0.00',
                uri: 'testing-event',
                file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
                business_id: '14161',
                business_address1: 'Lapangan D Senayan',
                country_name: 'Indonesia',
                city_name: 'Jakarta Selatan',
                voucher_provider: 'tiket.com',
                event_type: 'B',
                tiket_barcode: '6546216472445',
                checkin_date: null,
                tiket_cust_name: 'Abdul Rahman',
                tiket_gender: 'm',
                tiket_no_hp: '+6285692223310',
                tiket_birth_date: '1998-02-13',
                customer_price: '1000.00',
                customer_currency: 'IDR',
                tiket_cust_id: '09110110101',
                sell_rate_price: '1000.00',
                tiket_seating: '',
                tiket_attend_date: '0000-00-00',
                is_installment: '0',
                payment_status: 'shoppingcart',
                event_start: '2012-02-01 15:00:00',
                event_end: '2018-04-30 00:00:00',
                business_lat: '-6.20928550000000000000',
                business_long: '106.83307890000003000000',
                contact_person: '26042317',
                tiket_detail: [
                  {
                    tiket_barcode: '6546216472445',
                    tiket_cust_name: 'Abdul Rahman',
                    tiket_cust_id: '09110110101',
                    tiket_seating: '',
                    tiket_gender: 'm'
                  }
                ],
                tiket_quantity: 1,
                detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
                displayed_price: 1000,
                tax: 0,
                event_category: 'event',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                type: 'event',
                item_charge: 2000,
                item_charge_idr: 2000,
                tax_and_charge: 2000
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: false,
            is_change_payment: false,
            type: false,
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            bank_image: '/images/ico_bca.png',
            important_information: [
              'Informasi Penting',
              'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
              'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
              'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
              'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
            ],
            token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
          },
          loading: false,
          loaded: true
        },
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {
      type: 'SET_POPUP',
      payload: {
        show: true,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      }
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 12,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_image: '/images/ico_bca.png',
          important_information: [
            'Informasi Penting',
            'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
            'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
            'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
            'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
          ],
          token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: true,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: '0',
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                total_price: 1000,
                quantity: '1',
                order_id: '39982221',
                order_detail_id: '58944361',
                order_detail_status: 'active',
                order_type: 'event',
                order_master_id: '4777',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                selling_price: '1000.00',
                order_expire_datetime: '2018-04-04 11:59:43',
                selling_currency: 'IDR',
                created_timestamp: '2018-04-04 09:46:47',
                order_name: 'Testing Event',
                order_name_detail: 'Test CC',
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                tiket_with_seating: '0',
                tiket_required_info: 'name,idcard',
                tiket_min_purchase: '1',
                tiket_max_purchase: '10',
                tiket_start_sell: '2018-03-01 00:00:00',
                tiket_end_sell: '2018-06-30 00:00:00',
                ext_source: 'native',
                ext_source_id: null,
                tiket_sell_price_netto_api: '0.00',
                tiket_id: '4777',
                tiket_total_allotment: '1000',
                tiket_markup_price_api: '0.00',
                tiket_subsidy_price_api: '0.00',
                uri: 'testing-event',
                file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
                business_id: '14161',
                business_address1: 'Lapangan D Senayan',
                country_name: 'Indonesia',
                city_name: 'Jakarta Selatan',
                voucher_provider: 'tiket.com',
                event_type: 'B',
                tiket_barcode: '6546216472445',
                checkin_date: null,
                tiket_cust_name: 'Abdul Rahman',
                tiket_gender: 'm',
                tiket_no_hp: '+6285692223310',
                tiket_birth_date: '1998-02-13',
                customer_price: '1000.00',
                customer_currency: 'IDR',
                tiket_cust_id: '09110110101',
                sell_rate_price: '1000.00',
                tiket_seating: '',
                tiket_attend_date: '0000-00-00',
                is_installment: '0',
                payment_status: 'shoppingcart',
                event_start: '2012-02-01 15:00:00',
                event_end: '2018-04-30 00:00:00',
                business_lat: '-6.20928550000000000000',
                business_long: '106.83307890000003000000',
                contact_person: '26042317',
                tiket_detail: [
                  {
                    tiket_barcode: '6546216472445',
                    tiket_cust_name: 'Abdul Rahman',
                    tiket_cust_id: '09110110101',
                    tiket_seating: '',
                    tiket_gender: 'm'
                  }
                ],
                tiket_quantity: 1,
                detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
                displayed_price: 1000,
                tax: 0,
                event_category: 'event',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                type: 'event',
                item_charge: 2000,
                item_charge_idr: 2000,
                tax_and_charge: 2000
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: false,
            is_change_payment: false,
            type: false,
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            bank_image: '/images/ico_bca.png',
            important_information: [
              'Informasi Penting',
              'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
              'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
              'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
              'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
            ],
            token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
          },
          loading: false,
          loaded: true
        },
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {
          show: true,
          header: '',
          footer: {
            type: 'a',
            key: null,
            ref: null,
            props: {className: 'btn block large', children: 'OK'},
            _owner: null
          },
          content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
        },
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'RESET_POPUP'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 12,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_image: '/images/ico_bca.png',
          important_information: [
            'Informasi Penting',
            'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
            'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
            'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
            'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
          ],
          token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: false,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: '0',
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                total_price: 1000,
                quantity: '1',
                order_id: '39982221',
                order_detail_id: '58944361',
                order_detail_status: 'active',
                order_type: 'event',
                order_master_id: '4777',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                selling_price: '1000.00',
                order_expire_datetime: '2018-04-04 11:59:43',
                selling_currency: 'IDR',
                created_timestamp: '2018-04-04 09:46:47',
                order_name: 'Testing Event',
                order_name_detail: 'Test CC',
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                tiket_with_seating: '0',
                tiket_required_info: 'name,idcard',
                tiket_min_purchase: '1',
                tiket_max_purchase: '10',
                tiket_start_sell: '2018-03-01 00:00:00',
                tiket_end_sell: '2018-06-30 00:00:00',
                ext_source: 'native',
                ext_source_id: null,
                tiket_sell_price_netto_api: '0.00',
                tiket_id: '4777',
                tiket_total_allotment: '1000',
                tiket_markup_price_api: '0.00',
                tiket_subsidy_price_api: '0.00',
                uri: 'testing-event',
                file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
                business_id: '14161',
                business_address1: 'Lapangan D Senayan',
                country_name: 'Indonesia',
                city_name: 'Jakarta Selatan',
                voucher_provider: 'tiket.com',
                event_type: 'B',
                tiket_barcode: '6546216472445',
                checkin_date: null,
                tiket_cust_name: 'Abdul Rahman',
                tiket_gender: 'm',
                tiket_no_hp: '+6285692223310',
                tiket_birth_date: '1998-02-13',
                customer_price: '1000.00',
                customer_currency: 'IDR',
                tiket_cust_id: '09110110101',
                sell_rate_price: '1000.00',
                tiket_seating: '',
                tiket_attend_date: '0000-00-00',
                is_installment: '0',
                payment_status: 'shoppingcart',
                event_start: '2012-02-01 15:00:00',
                event_end: '2018-04-30 00:00:00',
                business_lat: '-6.20928550000000000000',
                business_long: '106.83307890000003000000',
                contact_person: '26042317',
                tiket_detail: [
                  {
                    tiket_barcode: '6546216472445',
                    tiket_cust_name: 'Abdul Rahman',
                    tiket_cust_id: '09110110101',
                    tiket_seating: '',
                    tiket_gender: 'm'
                  }
                ],
                tiket_quantity: 1,
                detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
                displayed_price: 1000,
                tax: 0,
                event_category: 'event',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                type: 'event',
                item_charge: 2000,
                item_charge_idr: 2000,
                tax_and_charge: 2000
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: false,
            is_change_payment: false,
            type: false,
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            bank_image: '/images/ico_bca.png',
            important_information: [
              'Informasi Penting',
              'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
              'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
              'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
              'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
            ],
            token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
          },
          loading: false,
          loaded: true
        },
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {
          show: false,
          header: '',
          footer: {
            type: 'a',
            key: null,
            ref: null,
            props: {className: 'btn block large', children: 'OK'},
            _owner: null
          },
          content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
        },
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_PAYMENT_CONFIRM'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 12,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          bank_image: '/images/ico_bca.png',
          important_information: [
            'Informasi Penting',
            'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
            'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
            'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
            'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
          ],
          token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: true, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: false,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: '0',
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                total_price: 1000,
                quantity: '1',
                order_id: '39982221',
                order_detail_id: '58944361',
                order_detail_status: 'active',
                order_type: 'event',
                order_master_id: '4777',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                selling_price: '1000.00',
                order_expire_datetime: '2018-04-04 11:59:43',
                selling_currency: 'IDR',
                created_timestamp: '2018-04-04 09:46:47',
                order_name: 'Testing Event',
                order_name_detail: 'Test CC',
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                tiket_with_seating: '0',
                tiket_required_info: 'name,idcard',
                tiket_min_purchase: '1',
                tiket_max_purchase: '10',
                tiket_start_sell: '2018-03-01 00:00:00',
                tiket_end_sell: '2018-06-30 00:00:00',
                ext_source: 'native',
                ext_source_id: null,
                tiket_sell_price_netto_api: '0.00',
                tiket_id: '4777',
                tiket_total_allotment: '1000',
                tiket_markup_price_api: '0.00',
                tiket_subsidy_price_api: '0.00',
                uri: 'testing-event',
                file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
                business_id: '14161',
                business_address1: 'Lapangan D Senayan',
                country_name: 'Indonesia',
                city_name: 'Jakarta Selatan',
                voucher_provider: 'tiket.com',
                event_type: 'B',
                tiket_barcode: '6546216472445',
                checkin_date: null,
                tiket_cust_name: 'Abdul Rahman',
                tiket_gender: 'm',
                tiket_no_hp: '+6285692223310',
                tiket_birth_date: '1998-02-13',
                customer_price: '1000.00',
                customer_currency: 'IDR',
                tiket_cust_id: '09110110101',
                sell_rate_price: '1000.00',
                tiket_seating: '',
                tiket_attend_date: '0000-00-00',
                is_installment: '0',
                payment_status: 'shoppingcart',
                event_start: '2012-02-01 15:00:00',
                event_end: '2018-04-30 00:00:00',
                business_lat: '-6.20928550000000000000',
                business_long: '106.83307890000003000000',
                contact_person: '26042317',
                tiket_detail: [
                  {
                    tiket_barcode: '6546216472445',
                    tiket_cust_name: 'Abdul Rahman',
                    tiket_cust_id: '09110110101',
                    tiket_seating: '',
                    tiket_gender: 'm'
                  }
                ],
                tiket_quantity: 1,
                detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
                displayed_price: 1000,
                tax: 0,
                event_category: 'event',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                type: 'event',
                item_charge: 2000,
                item_charge_idr: 2000,
                tax_and_charge: 2000
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: false,
            is_change_payment: false,
            type: false,
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            bank_image: '/images/ico_bca.png',
            important_information: [
              'Informasi Penting',
              'Pastikan jumlah dana yang anda transfer sesuai dengan Total Akhir  IDR 10.001,00  yang tertera di bawah',
              'Masukan ID pemesan anda dalam kolom berita transfer & simpan bukti pembayaran anda',
              'Transaksi akan dibatalkan (berakhir) jika anda tidak melakukan pembayaran pada periode waktu yang telah ditentukan atau nominal yang ditransfer tidak sesuai dengan total pembayaran',
              'E-Tiket atau Voucher akan dikirimkan melalui email setelah pembayaran selesai dilakukan'
            ],
            token: '7df071fa9dcf2b248ee4aa6305362c8c60962120'
          },
          loading: false,
          loaded: true
        },
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: true, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {
          show: false,
          header: '',
          footer: {
            type: 'a',
            key: null,
            ref: null,
            props: {className: 'btn block large', children: 'OK'},
            _owner: null
          },
          content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
        },
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'RESET_PAYMENT_DETAIL'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: true, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: false,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: true, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {
          show: false,
          header: '',
          footer: {
            type: 'a',
            key: null,
            ref: null,
            props: {className: 'btn block large', children: 'OK'},
            _owner: null
          },
          content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
        },
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_ACCOUNT'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: true, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: false,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: true, loaded: false},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: true, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {
          show: false,
          header: '',
          footer: {
            type: 'a',
            key: null,
            ref: null,
            props: {className: 'btn block large', children: 'OK'},
            _owner: null
          },
          content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
        },
        account: {data: {}, loading: true, loaded: false},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {result: {}, type: 'LOAD_ACCOUNT_SUCCESS'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: true, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: false,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: true, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {
          show: false,
          header: '',
          footer: {
            type: 'a',
            key: null,
            ref: null,
            props: {className: 'btn block large', children: 'OK'},
            _owner: null
          },
          content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
        },
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {
      result: {
        output_type: 'json',
        diagnostic: {
          status: 200,
          elapsetime: '0.7037',
          memoryusage: '36.21MB',
          unix_timestamp: 1522815517,
          confirm: 'success',
          lang: 'id',
          currency: 'IDR'
        },
        orderId: '39982221',
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: null,
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              quantity: '1',
              order_detail_id: '58944361',
              order_type: 'event',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              event_type: 'B'
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 12,
          is_confirmation: true,
          is_change_payment: false,
          type: '',
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          already_transfer_url: 'https://renan.tiket.com/confirmpayment/index/39982221/d8337c4e8ed29bf7494a0a37deebb39d'
        },
        banks: {
          bank_image: '/images/ico_bca.png',
          bank_owner_label: 'Nama',
          bank_owner: 'PT. Global Tiket Network',
          bank_name_label: 'Bank',
          bank_name: 'BCA',
          bank_branch_label: 'Cabang',
          bank_branch: 'Jakarta',
          bank_account_label: 'No Rekening',
          bank_account: '52 6032 2488'
        },
        message: 'Silakan transfer pembayaran untuk memesan. Anda memiliki waktu <strong> 26 menit  </strong> untuk melakukan pembayaran',
        grand_total: 10001,
        login_status: 'false',
        token: '16c01aefcd24df5ad8a91cce347f537c735eee69'
      },
      type: 'LOAD_PAYMENT_CONFIRM_SUCCESS'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {
        data: {
          output_type: 'json',
          diagnostic: {
            status: 200,
            elapsetime: '0.7037',
            memoryusage: '36.21MB',
            unix_timestamp: 1522815517,
            confirm: 'success',
            lang: 'id',
            currency: 'IDR'
          },
          orderId: '39982221',
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: null,
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                quantity: '1',
                order_detail_id: '58944361',
                order_type: 'event',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                event_type: 'B'
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: true,
            is_change_payment: false,
            type: '',
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            already_transfer_url: 'https://renan.tiket.com/confirmpayment/index/39982221/d8337c4e8ed29bf7494a0a37deebb39d'
          },
          banks: {
            bank_image: '/images/ico_bca.png',
            bank_owner_label: 'Nama',
            bank_owner: 'PT. Global Tiket Network',
            bank_name_label: 'Bank',
            bank_name: 'BCA',
            bank_branch_label: 'Cabang',
            bank_branch: 'Jakarta',
            bank_account_label: 'No Rekening',
            bank_account: '52 6032 2488'
          },
          message: 'Silakan transfer pembayaran untuk memesan. Anda memiliki waktu <strong> 26 menit  </strong> untuk melakukan pembayaran',
          grand_total: 10001,
          login_status: 'false',
          token: '16c01aefcd24df5ad8a91cce347f537c735eee69'
        },
        loading: false,
        loaded: true
      },
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: false, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {
        show: false,
        header: '',
        footer: {
          type: 'a',
          key: null,
          ref: null,
          props: {className: 'btn block large', children: 'OK'},
          _owner: null
        },
        content: 'Pastikan total transaksi Anda memenuhi ketentuan minimum Transaksi Gift Voucher.'
      },
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {result: {}, loading: true, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: false, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: false, loaded: false, data: {}},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'LOAD_ORDER_DETAIL'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {result: {}, loading: true, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {result: {}, loading: true, loaded: false},
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {loading: false, loaded: false, data: {}},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
        paymentDetail: {
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: '0',
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                total_price: 1000,
                quantity: '1',
                order_id: '39982221',
                order_detail_id: '58944361',
                order_detail_status: 'active',
                order_type: 'event',
                order_master_id: '4777',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                selling_price: '1000.00',
                order_expire_datetime: '2018-04-04 11:59:43',
                selling_currency: 'IDR',
                created_timestamp: '2018-04-04 09:46:47',
                order_name: 'Testing Event',
                order_name_detail: 'Test CC',
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                tiket_with_seating: '0',
                tiket_required_info: 'name,idcard',
                tiket_min_purchase: '1',
                tiket_max_purchase: '10',
                tiket_start_sell: '2018-03-01 00:00:00',
                tiket_end_sell: '2018-06-30 00:00:00',
                ext_source: 'native',
                ext_source_id: null,
                tiket_sell_price_netto_api: '0.00',
                tiket_id: '4777',
                tiket_total_allotment: '1000',
                tiket_markup_price_api: '0.00',
                tiket_subsidy_price_api: '0.00',
                uri: 'testing-event',
                file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
                business_id: '14161',
                business_address1: 'Lapangan D Senayan',
                country_name: 'Indonesia',
                city_name: 'Jakarta Selatan',
                voucher_provider: 'tiket.com',
                event_type: 'B',
                tiket_barcode: '6546216472445',
                checkin_date: null,
                tiket_cust_name: 'Abdul Rahman',
                tiket_gender: 'm',
                tiket_no_hp: '+6285692223310',
                tiket_birth_date: '1998-02-13',
                customer_price: '1000.00',
                customer_currency: 'IDR',
                tiket_cust_id: '09110110101',
                sell_rate_price: '1000.00',
                tiket_seating: '',
                tiket_attend_date: '0000-00-00',
                is_installment: '0',
                payment_status: 'shoppingcart',
                event_start: '2012-02-01 15:00:00',
                event_end: '2018-04-30 00:00:00',
                business_lat: '-6.20928550000000000000',
                business_long: '106.83307890000003000000',
                contact_person: '26042317',
                tiket_detail: [
                  {
                    tiket_barcode: '6546216472445',
                    tiket_cust_name: 'Abdul Rahman',
                    tiket_cust_id: '09110110101',
                    tiket_seating: '',
                    tiket_gender: 'm'
                  }
                ],
                tiket_quantity: 1,
                detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
                displayed_price: 1000,
                tax: 0,
                event_category: 'event',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                type: 'event',
                item_charge: 2000,
                item_charge_idr: 2000,
                tax_and_charge: 2000
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 3,
            is_confirmation: false,
            is_change_payment: false,
            type: false,
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            token: '8a53cfe7d3364352c082aca1377ccc92e838b92a'
          },
          loading: false,
          loaded: true
        },
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {data: {}, loading: false, loaded: false},
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {result: {}, loading: true, loaded: false},
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {loading: false, loaded: false, data: {}},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {
      result: {
        output_type: 'json',
        myorder: {
          order_id: '39982221',
          data: [
            {
              expire: 19,
              order_detail_id: '58944361',
              order_expire_datetime: '2018-04-04 11:44:43',
              order_type: 'event',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              order_detail_status: 'active',
              detail: {
                order_detail_id: '58944361',
                tiket_barcode: '6546216472445',
                qty: '1',
                sell_rate_price: '1000.00',
                sell_rate_currency: 'IDR',
                tiket_seating: '',
                startdate: '2012-02-01 00:00:00',
                enddate: '2018-06-30 00:00:00',
                event_type: 'B',
                event_category: 'event',
                event_address: 'Lapangan D Senayan',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                attend_date: null
              },
              order_photo: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              order_icon: 'h3d',
              tax_and_charge: '5000.00',
              subtotal_and_charge: '6000.00',
              delete_uri: 'https://renan.tiket.com/order/delete_order?order_detail_id=58944361&order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              business_id: '14161'
            }
          ],
          total_IDR: 1000,
          total: 6000,
          total_tax: 5000,
          total_without_tax: 1000,
          count_installment: 0,
          breakdown_price: [{title: '1 x Test CC', value: 1000}],
          promo: [],
          discount: 0,
          discount_amount: 0,
          total_without_promocode: 1000,
          payment_status: 'shoppingcart'
        },
        diagnostic: {
          status: 200,
          elapsetime: '1.4200',
          memoryusage: '38.78MB',
          unix_timestamp: 1522815903,
          confirm: 'success',
          lang: 'id',
          currency: 'IDR'
        },
        checkout: 'https://renan.tiket.com/order/checkout/39982221/IDR',
        login_status: 'false',
        token: '5021081ba0bd4b43953414f0bfb0f03b61b4943b'
      },
      type: 'LOAD_ORDER_DETAIL_SUCCESS'
    }
  );
  expect(state).toEqual({
    payment: {
      paymentList: {paymentMethods: [], paymentDescription: {}, loading: false, loaded: false},
      paymentDetail: {
        result: {
          payment_subsider_tiket: 0,
          currency_to_be_converted: 'IDR',
          from_another_currency: false,
          reseller_id: '0',
          payment_discount: 0,
          order_types: ['event'],
          order_id: '39982221',
          order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          currency: 'IDR',
          payment_charge: 2000,
          giftPromo: false,
          sub_total: 1000,
          unique_code: 7001,
          grand_total: 10001,
          grand_subtotal: 3000,
          orders: [
            {
              total_price: 1000,
              quantity: '1',
              order_id: '39982221',
              order_detail_id: '58944361',
              order_detail_status: 'active',
              order_type: 'event',
              order_master_id: '4777',
              event_name: 'Testing Event',
              tiket_name: 'Test CC',
              currency: 'IDR',
              price: 1000,
              selling_price: '1000.00',
              order_expire_datetime: '2018-04-04 11:59:43',
              selling_currency: 'IDR',
              created_timestamp: '2018-04-04 09:46:47',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              tiket_event_start: '2012-02-01 00:00:00',
              tiket_event_end: '2018-06-30 00:00:00',
              tiket_with_seating: '0',
              tiket_required_info: 'name,idcard',
              tiket_min_purchase: '1',
              tiket_max_purchase: '10',
              tiket_start_sell: '2018-03-01 00:00:00',
              tiket_end_sell: '2018-06-30 00:00:00',
              ext_source: 'native',
              ext_source_id: null,
              tiket_sell_price_netto_api: '0.00',
              tiket_id: '4777',
              tiket_total_allotment: '1000',
              tiket_markup_price_api: '0.00',
              tiket_subsidy_price_api: '0.00',
              uri: 'testing-event',
              file_name: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              business_id: '14161',
              business_address1: 'Lapangan D Senayan',
              country_name: 'Indonesia',
              city_name: 'Jakarta Selatan',
              voucher_provider: 'tiket.com',
              event_type: 'B',
              tiket_barcode: '6546216472445',
              checkin_date: null,
              tiket_cust_name: 'Abdul Rahman',
              tiket_gender: 'm',
              tiket_no_hp: '+6285692223310',
              tiket_birth_date: '1998-02-13',
              customer_price: '1000.00',
              customer_currency: 'IDR',
              tiket_cust_id: '09110110101',
              sell_rate_price: '1000.00',
              tiket_seating: '',
              tiket_attend_date: '0000-00-00',
              is_installment: '0',
              payment_status: 'shoppingcart',
              event_start: '2012-02-01 15:00:00',
              event_end: '2018-04-30 00:00:00',
              business_lat: '-6.20928550000000000000',
              business_long: '106.83307890000003000000',
              contact_person: '26042317',
              tiket_detail: [
                {
                  tiket_barcode: '6546216472445',
                  tiket_cust_name: 'Abdul Rahman',
                  tiket_cust_id: '09110110101',
                  tiket_seating: '',
                  tiket_gender: 'm'
                }
              ],
              tiket_quantity: 1,
              detail_ticket_schedule: '01 Feb 2012 - 30 Jun 2018',
              displayed_price: 1000,
              tax: 0,
              event_category: 'event',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              type: 'event',
              item_charge: 2000,
              item_charge_idr: 2000,
              tax_and_charge: 2000
            }
          ],
          confirm_page_mobile: false,
          gaq: '',
          payment_type: 3,
          is_confirmation: false,
          is_change_payment: false,
          type: false,
          checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          arrOrderType: ['event'],
          order_expire_datetime: '2018-04-04 11:44:43',
          tiket_point: 0,
          tiket_point_worth: 0,
          tiket_point_notes: '',
          tiket_point_status: '',
          tiket_point_words: '',
          token: '8a53cfe7d3364352c082aca1377ccc92e838b92a'
        },
        loading: false,
        loaded: true
      },
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {data: {}, loading: false, loaded: false},
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {
      result: {
        order_id: '39982221',
        data: [
          {
            expire: 19,
            order_detail_id: '58944361',
            order_expire_datetime: '2018-04-04 11:44:43',
            order_type: 'event',
            order_name: 'Testing Event',
            order_name_detail: 'Test CC',
            order_detail_status: 'active',
            detail: {
              order_detail_id: '58944361',
              tiket_barcode: '6546216472445',
              qty: '1',
              sell_rate_price: '1000.00',
              sell_rate_currency: 'IDR',
              tiket_seating: '',
              startdate: '2012-02-01 00:00:00',
              enddate: '2018-06-30 00:00:00',
              event_type: 'B',
              event_category: 'event',
              event_address: 'Lapangan D Senayan',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              attend_date: null
            },
            order_photo: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
            order_icon: 'h3d',
            tax_and_charge: '5000.00',
            subtotal_and_charge: '6000.00',
            delete_uri: 'https://renan.tiket.com/order/delete_order?order_detail_id=58944361&order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            business_id: '14161'
          }
        ],
        total_IDR: 1000,
        total: 6000,
        total_tax: 5000,
        total_without_tax: 1000,
        count_installment: 0,
        breakdown_price: [{title: '1 x Test CC', value: 1000}],
        promo: [],
        discount: 0,
        discount_amount: 0,
        total_without_promocode: 1000,
        payment_status: 'shoppingcart'
      },
      loading: false,
      loaded: true
    },
    app: {
      flash: {show: false, type: '', text: ''},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {loading: false, loaded: false, data: {}},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {
          paymentMethods: [
            {
              code: '1',
              link: 'https://renan.tiket.com/checkout/checkout_payment/1?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Kartu Kredit',
              message: '',
              type: 'creditcard',
              desc: 'Master Card and Visa',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/visa.png',
                'https://renan.tiket.com/images/apps_payment/master_card.png'
              ],
              countdown_time: 1061
            },
            {
              code: '2',
              link: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Transfer',
              message: '',
              type: 'banktransfer',
              desc: 'Transfer',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png'
              ],
              payment_group: [
                {
                  code: '12',
                  type: 'bca_transfer',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'BCA Transfer',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                  desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                },
                {
                  code: '42',
                  type: 'mandiri_transfer',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'Mandiri Transfer',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                  desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                }
              ],
              countdown_time: 1061
            },
            {
              code: '99',
              type: 'virtualaccount',
              link: 'https://renan.tiket.com/checkout/checkout_payment/99?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Virtual Account',
              message: '',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                'https://renan.tiket.com/images/apps_payment/bri_transfer.png'
              ],
              desc: 'Virtual Account',
              payment_group: [
                {
                  code: '13',
                  type: 'va_bca',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/13?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA BCA',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                },
                {
                  code: '43',
                  type: 'va_mandiri',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/43?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA Mandiri',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                },
                {
                  code: '52',
                  type: 'va_bni',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/52?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA BNI',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking, Mobile Banking & SMS Banking',
                  countdown_time: 1061
                },
                {
                  code: '62',
                  type: 'va_bri',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/62?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA BRI',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bri_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                }
              ],
              countdown_time: 1061
            },
            {
              code: '59',
              link: 'https://renan.tiket.com/checkout/checkout_payment/59?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'ATM',
              message: '',
              type: 'atm_nicepay',
              desc: 'ATM Bersama, Prima & Alto',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/atm_bersama.png',
                'https://renan.tiket.com/images/apps_payment/atm_prima.png',
                'https://renan.tiket.com/images/apps_payment/alto.png'
              ],
              countdown_time: 1061
            },
            {
              code: '3',
              link: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'KlikBCA',
              message: '',
              type: 'klikbca',
              desc: 'KlikBCA',
              app_images: ['https://renan.tiket.com/images/apps_payment/klik_bca.png'],
              countdown_time: 1061
            },
            {
              code: '4',
              link: 'https://renan.tiket.com/checkout/checkout_payment/4?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'BCA KlikPay',
              message: '',
              type: 'klikpay',
              desc: 'BCA Klikpay',
              app_images: ['https://renan.tiket.com/images/apps_payment/bca_klikpay.png'],
              countdown_time: 1061
            },
            {
              code: '34',
              link: 'https://renan.tiket.com/checkout/checkout_payment/34?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Mandiri Clickpay',
              message: 'Butuh No. Kartu dan Token',
              type: 'mandiri_clickpay',
              desc: 'Mandiri Clickpay',
              app_images: ['https://renan.tiket.com/images/apps_payment/mandiri_clickpay.png'],
              countdown_time: 1061
            },
            {
              code: '31',
              link: 'https://renan.tiket.com/checkout/checkout_payment/31?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'CIMB Clicks',
              message: '',
              type: 'cimbclicks',
              desc: 'CIMB Clicks',
              app_images: ['https://renan.tiket.com/images/apps_payment/cimb_click.png'],
              countdown_time: 1061
            },
            {
              code: '33',
              link: 'https://renan.tiket.com/checkout/checkout_payment/33?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'ePay BRI',
              message: '',
              type: 'epaybri',
              desc: 'ePay BRI',
              app_images: ['https://renan.tiket.com/images/apps_payment/e_pay_bri.png'],
              countdown_time: 1061
            }
          ],
          paymentDescription: {
            currency: 'IDR',
            orderId: '39982221',
            orderHash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            grandTotal: 10001,
            subTotal: 1000,
            grandSubTotal: 3000,
            paymentCharge: 2000,
            baggageFee: [0],
            uniqueCode: 7001,
            orderExpiredDatetime: '2018-04-04 11:44:43',
            paymentType: 3,
            tixPoint: 0,
            giftPromo: false
          },
          loading: false,
          loaded: true
        },
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {
          data: {
            output_type: 'json',
            diagnostic: {
              status: 200,
              elapsetime: '0.7447',
              memoryusage: '38.13MB',
              unix_timestamp: 1522816063,
              confirm: 'success',
              lang: 'id',
              currency: 'IDR'
            },
            orderId: '39982221',
            result: {
              payment_subsider_tiket: 0,
              currency_to_be_converted: 'IDR',
              from_another_currency: false,
              reseller_id: null,
              payment_discount: 0,
              order_types: ['event'],
              order_id: '39982221',
              order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              currency: 'IDR',
              payment_charge: 2000,
              giftPromo: false,
              sub_total: 1000,
              unique_code: 7001,
              grand_total: 10001,
              grand_subtotal: 3000,
              orders: [
                {
                  quantity: '1',
                  order_detail_id: '58944361',
                  order_type: 'event',
                  event_name: 'Testing Event',
                  tiket_name: 'Test CC',
                  currency: 'IDR',
                  price: 1000,
                  tiket_event_start: '2012-02-01 00:00:00',
                  tiket_event_end: '2018-06-30 00:00:00',
                  event_type: 'B'
                }
              ],
              confirm_page_mobile: false,
              gaq: '',
              payment_type: 12,
              is_confirmation: true,
              is_change_payment: false,
              type: '',
              checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              arrOrderType: ['event'],
              order_expire_datetime: '2018-04-04 11:44:43',
              tiket_point: 0,
              tiket_point_worth: 0,
              tiket_point_notes: '',
              tiket_point_status: '',
              tiket_point_words: '',
              already_transfer_url: 'https://renan.tiket.com/confirmpayment/index/39982221/d8337c4e8ed29bf7494a0a37deebb39d'
            },
            banks: {
              bank_image: '/images/ico_bca.png',
              bank_owner_label: 'Nama',
              bank_owner: 'PT. Global Tiket Network',
              bank_name_label: 'Bank',
              bank_name: 'BCA',
              bank_branch_label: 'Cabang',
              bank_branch: 'Jakarta',
              bank_account_label: 'No Rekening',
              bank_account: '52 6032 2488'
            },
            message: 'Silakan transfer pembayaran untuk memesan. Anda memiliki waktu <strong> 17 menit  </strong> untuk melakukan pembayaran',
            grand_total: 10001,
            login_status: 'false',
            token: 'a5d579bfd821fa4bbba9c88a5fdf3f320b2affb7'
          },
          loading: false,
          loaded: true
        },
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {
        result: {
          order_id: '39982221',
          data: [
            {
              expire: 19,
              order_detail_id: '58944361',
              order_expire_datetime: '2018-04-04 11:44:43',
              order_type: 'event',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              order_detail_status: 'active',
              detail: {
                order_detail_id: '58944361',
                tiket_barcode: '6546216472445',
                qty: '1',
                sell_rate_price: '1000.00',
                sell_rate_currency: 'IDR',
                tiket_seating: '',
                startdate: '2012-02-01 00:00:00',
                enddate: '2018-06-30 00:00:00',
                event_type: 'B',
                event_category: 'event',
                event_address: 'Lapangan D Senayan',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                attend_date: null
              },
              order_photo: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              order_icon: 'h3d',
              tax_and_charge: '5000.00',
              subtotal_and_charge: '6000.00',
              delete_uri: 'https://renan.tiket.com/order/delete_order?order_detail_id=58944361&order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              business_id: '14161'
            }
          ],
          total_IDR: 1000,
          total: 6000,
          total_tax: 5000,
          total_without_tax: 1000,
          count_installment: 0,
          breakdown_price: [{title: '1 x Test CC', value: 1000}],
          promo: [],
          discount: 0,
          discount_amount: 0,
          total_without_promocode: 1000,
          payment_status: 'shoppingcart'
        },
        loading: false,
        loaded: true
      },
      app: {
        flash: {show: false, type: '', text: ''},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'SET_FLASH', payload: {show: true, type: 'success', text: 'Total Payment Copied.'}}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {
        paymentMethods: [
          {
            code: '1',
            link: 'https://renan.tiket.com/checkout/checkout_payment/1?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Kartu Kredit',
            message: '',
            type: 'creditcard',
            desc: 'Master Card and Visa',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/visa.png',
              'https://renan.tiket.com/images/apps_payment/master_card.png'
            ],
            countdown_time: 1061
          },
          {
            code: '2',
            link: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Transfer',
            message: '',
            type: 'banktransfer',
            desc: 'Transfer',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png'
            ],
            payment_group: [
              {
                code: '12',
                type: 'bca_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'BCA Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              },
              {
                code: '42',
                type: 'mandiri_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'Mandiri Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              }
            ],
            countdown_time: 1061
          },
          {
            code: '99',
            type: 'virtualaccount',
            link: 'https://renan.tiket.com/checkout/checkout_payment/99?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Virtual Account',
            message: '',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bri_transfer.png'
            ],
            desc: 'Virtual Account',
            payment_group: [
              {
                code: '13',
                type: 'va_bca',
                link: 'http://renan.tiket.com/checkout/checkout_payment/13?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA BCA',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              },
              {
                code: '43',
                type: 'va_mandiri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/43?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA Mandiri',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              },
              {
                code: '52',
                type: 'va_bni',
                link: 'http://renan.tiket.com/checkout/checkout_payment/52?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA BNI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking, Mobile Banking & SMS Banking',
                countdown_time: 1061
              },
              {
                code: '62',
                type: 'va_bri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/62?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA BRI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              }
            ],
            countdown_time: 1061
          },
          {
            code: '59',
            link: 'https://renan.tiket.com/checkout/checkout_payment/59?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'ATM',
            message: '',
            type: 'atm_nicepay',
            desc: 'ATM Bersama, Prima & Alto',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/atm_bersama.png',
              'https://renan.tiket.com/images/apps_payment/atm_prima.png',
              'https://renan.tiket.com/images/apps_payment/alto.png'
            ],
            countdown_time: 1061
          },
          {
            code: '3',
            link: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'KlikBCA',
            message: '',
            type: 'klikbca',
            desc: 'KlikBCA',
            app_images: ['https://renan.tiket.com/images/apps_payment/klik_bca.png'],
            countdown_time: 1061
          },
          {
            code: '4',
            link: 'https://renan.tiket.com/checkout/checkout_payment/4?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'BCA KlikPay',
            message: '',
            type: 'klikpay',
            desc: 'BCA Klikpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/bca_klikpay.png'],
            countdown_time: 1061
          },
          {
            code: '34',
            link: 'https://renan.tiket.com/checkout/checkout_payment/34?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Mandiri Clickpay',
            message: 'Butuh No. Kartu dan Token',
            type: 'mandiri_clickpay',
            desc: 'Mandiri Clickpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/mandiri_clickpay.png'],
            countdown_time: 1061
          },
          {
            code: '31',
            link: 'https://renan.tiket.com/checkout/checkout_payment/31?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'CIMB Clicks',
            message: '',
            type: 'cimbclicks',
            desc: 'CIMB Clicks',
            app_images: ['https://renan.tiket.com/images/apps_payment/cimb_click.png'],
            countdown_time: 1061
          },
          {
            code: '33',
            link: 'https://renan.tiket.com/checkout/checkout_payment/33?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'ePay BRI',
            message: '',
            type: 'epaybri',
            desc: 'ePay BRI',
            app_images: ['https://renan.tiket.com/images/apps_payment/e_pay_bri.png'],
            countdown_time: 1061
          }
        ],
        paymentDescription: {
          currency: 'IDR',
          orderId: '39982221',
          orderHash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          grandTotal: 10001,
          subTotal: 1000,
          grandSubTotal: 3000,
          paymentCharge: 2000,
          baggageFee: [0],
          uniqueCode: 7001,
          orderExpiredDatetime: '2018-04-04 11:44:43',
          paymentType: 3,
          tixPoint: 0,
          giftPromo: false
        },
        loading: false,
        loaded: true
      },
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {
        data: {
          output_type: 'json',
          diagnostic: {
            status: 200,
            elapsetime: '0.7447',
            memoryusage: '38.13MB',
            unix_timestamp: 1522816063,
            confirm: 'success',
            lang: 'id',
            currency: 'IDR'
          },
          orderId: '39982221',
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: null,
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                quantity: '1',
                order_detail_id: '58944361',
                order_type: 'event',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                event_type: 'B'
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: true,
            is_change_payment: false,
            type: '',
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            already_transfer_url: 'https://renan.tiket.com/confirmpayment/index/39982221/d8337c4e8ed29bf7494a0a37deebb39d'
          },
          banks: {
            bank_image: '/images/ico_bca.png',
            bank_owner_label: 'Nama',
            bank_owner: 'PT. Global Tiket Network',
            bank_name_label: 'Bank',
            bank_name: 'BCA',
            bank_branch_label: 'Cabang',
            bank_branch: 'Jakarta',
            bank_account_label: 'No Rekening',
            bank_account: '52 6032 2488'
          },
          message: 'Silakan transfer pembayaran untuk memesan. Anda memiliki waktu <strong> 17 menit  </strong> untuk melakukan pembayaran',
          grand_total: 10001,
          login_status: 'false',
          token: 'a5d579bfd821fa4bbba9c88a5fdf3f320b2affb7'
        },
        loading: false,
        loaded: true
      },
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {
      result: {
        order_id: '39982221',
        data: [
          {
            expire: 19,
            order_detail_id: '58944361',
            order_expire_datetime: '2018-04-04 11:44:43',
            order_type: 'event',
            order_name: 'Testing Event',
            order_name_detail: 'Test CC',
            order_detail_status: 'active',
            detail: {
              order_detail_id: '58944361',
              tiket_barcode: '6546216472445',
              qty: '1',
              sell_rate_price: '1000.00',
              sell_rate_currency: 'IDR',
              tiket_seating: '',
              startdate: '2012-02-01 00:00:00',
              enddate: '2018-06-30 00:00:00',
              event_type: 'B',
              event_category: 'event',
              event_address: 'Lapangan D Senayan',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              attend_date: null
            },
            order_photo: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
            order_icon: 'h3d',
            tax_and_charge: '5000.00',
            subtotal_and_charge: '6000.00',
            delete_uri: 'https://renan.tiket.com/order/delete_order?order_detail_id=58944361&order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            business_id: '14161'
          }
        ],
        total_IDR: 1000,
        total: 6000,
        total_tax: 5000,
        total_without_tax: 1000,
        count_installment: 0,
        breakdown_price: [{title: '1 x Test CC', value: 1000}],
        promo: [],
        discount: 0,
        discount_amount: 0,
        total_without_promocode: 1000,
        payment_status: 'shoppingcart'
      },
      loading: false,
      loaded: true
    },
    app: {
      flash: {show: true, type: 'success', text: 'Total Payment Copied.'},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
  state = reducers(
    {
      payment: {
        paymentList: {
          paymentMethods: [
            {
              code: '1',
              link: 'https://renan.tiket.com/checkout/checkout_payment/1?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Kartu Kredit',
              message: '',
              type: 'creditcard',
              desc: 'Master Card and Visa',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/visa.png',
                'https://renan.tiket.com/images/apps_payment/master_card.png'
              ],
              countdown_time: 1061
            },
            {
              code: '2',
              link: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Transfer',
              message: '',
              type: 'banktransfer',
              desc: 'Transfer',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png'
              ],
              payment_group: [
                {
                  code: '12',
                  type: 'bca_transfer',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'BCA Transfer',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                  desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                },
                {
                  code: '42',
                  type: 'mandiri_transfer',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'Mandiri Transfer',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                  desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                }
              ],
              countdown_time: 1061
            },
            {
              code: '99',
              type: 'virtualaccount',
              link: 'https://renan.tiket.com/checkout/checkout_payment/99?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Virtual Account',
              message: '',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                'https://renan.tiket.com/images/apps_payment/bri_transfer.png'
              ],
              desc: 'Virtual Account',
              payment_group: [
                {
                  code: '13',
                  type: 'va_bca',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/13?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA BCA',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                },
                {
                  code: '43',
                  type: 'va_mandiri',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/43?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA Mandiri',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                },
                {
                  code: '52',
                  type: 'va_bni',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/52?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA BNI',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking, Mobile Banking & SMS Banking',
                  countdown_time: 1061
                },
                {
                  code: '62',
                  type: 'va_bri',
                  link: 'http://renan.tiket.com/checkout/checkout_payment/62?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                  text: 'VA BRI',
                  message: '',
                  images: 'https://renan.tiket.com/images/apps_payment/bri_transfer.png',
                  desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                  countdown_time: 1061
                }
              ],
              countdown_time: 1061
            },
            {
              code: '59',
              link: 'https://renan.tiket.com/checkout/checkout_payment/59?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'ATM',
              message: '',
              type: 'atm_nicepay',
              desc: 'ATM Bersama, Prima & Alto',
              app_images: [
                'https://renan.tiket.com/images/apps_payment/atm_bersama.png',
                'https://renan.tiket.com/images/apps_payment/atm_prima.png',
                'https://renan.tiket.com/images/apps_payment/alto.png'
              ],
              countdown_time: 1061
            },
            {
              code: '3',
              link: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'KlikBCA',
              message: '',
              type: 'klikbca',
              desc: 'KlikBCA',
              app_images: ['https://renan.tiket.com/images/apps_payment/klik_bca.png'],
              countdown_time: 1061
            },
            {
              code: '4',
              link: 'https://renan.tiket.com/checkout/checkout_payment/4?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'BCA KlikPay',
              message: '',
              type: 'klikpay',
              desc: 'BCA Klikpay',
              app_images: ['https://renan.tiket.com/images/apps_payment/bca_klikpay.png'],
              countdown_time: 1061
            },
            {
              code: '34',
              link: 'https://renan.tiket.com/checkout/checkout_payment/34?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'Mandiri Clickpay',
              message: 'Butuh No. Kartu dan Token',
              type: 'mandiri_clickpay',
              desc: 'Mandiri Clickpay',
              app_images: ['https://renan.tiket.com/images/apps_payment/mandiri_clickpay.png'],
              countdown_time: 1061
            },
            {
              code: '31',
              link: 'https://renan.tiket.com/checkout/checkout_payment/31?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'CIMB Clicks',
              message: '',
              type: 'cimbclicks',
              desc: 'CIMB Clicks',
              app_images: ['https://renan.tiket.com/images/apps_payment/cimb_click.png'],
              countdown_time: 1061
            },
            {
              code: '33',
              link: 'https://renan.tiket.com/checkout/checkout_payment/33?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              text: 'ePay BRI',
              message: '',
              type: 'epaybri',
              desc: 'ePay BRI',
              app_images: ['https://renan.tiket.com/images/apps_payment/e_pay_bri.png'],
              countdown_time: 1061
            }
          ],
          paymentDescription: {
            currency: 'IDR',
            orderId: '39982221',
            orderHash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            grandTotal: 10001,
            subTotal: 1000,
            grandSubTotal: 3000,
            paymentCharge: 2000,
            baggageFee: [0],
            uniqueCode: 7001,
            orderExpiredDatetime: '2018-04-04 11:44:43',
            paymentType: 3,
            tixPoint: 0,
            giftPromo: false
          },
          loading: false,
          loaded: true
        },
        paymentDetail: {result: {}, loading: false, loaded: false},
        paymentGroup: {result: {}, loading: false, loaded: false},
        paymentConfirm: {
          data: {
            output_type: 'json',
            diagnostic: {
              status: 200,
              elapsetime: '0.7447',
              memoryusage: '38.13MB',
              unix_timestamp: 1522816063,
              confirm: 'success',
              lang: 'id',
              currency: 'IDR'
            },
            orderId: '39982221',
            result: {
              payment_subsider_tiket: 0,
              currency_to_be_converted: 'IDR',
              from_another_currency: false,
              reseller_id: null,
              payment_discount: 0,
              order_types: ['event'],
              order_id: '39982221',
              order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              currency: 'IDR',
              payment_charge: 2000,
              giftPromo: false,
              sub_total: 1000,
              unique_code: 7001,
              grand_total: 10001,
              grand_subtotal: 3000,
              orders: [
                {
                  quantity: '1',
                  order_detail_id: '58944361',
                  order_type: 'event',
                  event_name: 'Testing Event',
                  tiket_name: 'Test CC',
                  currency: 'IDR',
                  price: 1000,
                  tiket_event_start: '2012-02-01 00:00:00',
                  tiket_event_end: '2018-06-30 00:00:00',
                  event_type: 'B'
                }
              ],
              confirm_page_mobile: false,
              gaq: '',
              payment_type: 12,
              is_confirmation: true,
              is_change_payment: false,
              type: '',
              checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              arrOrderType: ['event'],
              order_expire_datetime: '2018-04-04 11:44:43',
              tiket_point: 0,
              tiket_point_worth: 0,
              tiket_point_notes: '',
              tiket_point_status: '',
              tiket_point_words: '',
              already_transfer_url: 'https://renan.tiket.com/confirmpayment/index/39982221/d8337c4e8ed29bf7494a0a37deebb39d'
            },
            banks: {
              bank_image: '/images/ico_bca.png',
              bank_owner_label: 'Nama',
              bank_owner: 'PT. Global Tiket Network',
              bank_name_label: 'Bank',
              bank_name: 'BCA',
              bank_branch_label: 'Cabang',
              bank_branch: 'Jakarta',
              bank_account_label: 'No Rekening',
              bank_account: '52 6032 2488'
            },
            message: 'Silakan transfer pembayaran untuk memesan. Anda memiliki waktu <strong> 17 menit  </strong> untuk melakukan pembayaran',
            grand_total: 10001,
            login_status: 'false',
            token: 'a5d579bfd821fa4bbba9c88a5fdf3f320b2affb7'
          },
          loading: false,
          loaded: true
        },
        paymentVerify: {data: {}, loading: false, loaded: false}
      },
      order: {
        result: {
          order_id: '39982221',
          data: [
            {
              expire: 19,
              order_detail_id: '58944361',
              order_expire_datetime: '2018-04-04 11:44:43',
              order_type: 'event',
              order_name: 'Testing Event',
              order_name_detail: 'Test CC',
              order_detail_status: 'active',
              detail: {
                order_detail_id: '58944361',
                tiket_barcode: '6546216472445',
                qty: '1',
                sell_rate_price: '1000.00',
                sell_rate_currency: 'IDR',
                tiket_seating: '',
                startdate: '2012-02-01 00:00:00',
                enddate: '2018-06-30 00:00:00',
                event_type: 'B',
                event_category: 'event',
                event_address: 'Lapangan D Senayan',
                travellers: [
                  {
                    full_name: 'Abdul Rahman',
                    id_card: '09110110101',
                    salutation: 'Tuan',
                    seat: null
                  }
                ],
                attend_date: null
              },
              order_photo: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
              order_icon: 'h3d',
              tax_and_charge: '5000.00',
              subtotal_and_charge: '6000.00',
              delete_uri: 'https://renan.tiket.com/order/delete_order?order_detail_id=58944361&order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
              business_id: '14161'
            }
          ],
          total_IDR: 1000,
          total: 6000,
          total_tax: 5000,
          total_without_tax: 1000,
          count_installment: 0,
          breakdown_price: [{title: '1 x Test CC', value: 1000}],
          promo: [],
          discount: 0,
          discount_amount: 0,
          total_without_promocode: 1000,
          payment_status: 'shoppingcart'
        },
        loading: false,
        loaded: true
      },
      app: {
        flash: {show: true, type: 'success', text: 'Total Payment Copied.'},
        popup: {show: false, header: '', footer: '', content: ''},
        account: {data: {}, loading: false, loaded: true},
        context: {
          query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
          params: {}
        }
      }
    },
    {type: 'RESET_FLASH'}
  );
  expect(state).toEqual({
    payment: {
      paymentList: {
        paymentMethods: [
          {
            code: '1',
            link: 'https://renan.tiket.com/checkout/checkout_payment/1?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Kartu Kredit',
            message: '',
            type: 'creditcard',
            desc: 'Master Card and Visa',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/visa.png',
              'https://renan.tiket.com/images/apps_payment/master_card.png'
            ],
            countdown_time: 1061
          },
          {
            code: '2',
            link: 'https://renan.tiket.com/checkout/checkout_payment/2?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Transfer',
            message: '',
            type: 'banktransfer',
            desc: 'Transfer',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png'
            ],
            payment_group: [
              {
                code: '12',
                type: 'bca_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'BCA Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              },
              {
                code: '42',
                type: 'mandiri_transfer',
                link: 'http://renan.tiket.com/checkout/checkout_payment/42?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'Mandiri Transfer',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini hanya untuk menerima transfer lewat ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              }
            ],
            countdown_time: 1061
          },
          {
            code: '99',
            type: 'virtualaccount',
            link: 'https://renan.tiket.com/checkout/checkout_payment/99?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Virtual Account',
            message: '',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
              'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
              'https://renan.tiket.com/images/apps_payment/bri_transfer.png'
            ],
            desc: 'Virtual Account',
            payment_group: [
              {
                code: '13',
                type: 'va_bca',
                link: 'http://renan.tiket.com/checkout/checkout_payment/13?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA BCA',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bca_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              },
              {
                code: '43',
                type: 'va_mandiri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/43?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA Mandiri',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/mandiri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              },
              {
                code: '52',
                type: 'va_bni',
                link: 'http://renan.tiket.com/checkout/checkout_payment/52?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA BNI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bni_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking, Mobile Banking & SMS Banking',
                countdown_time: 1061
              },
              {
                code: '62',
                type: 'va_bri',
                link: 'http://renan.tiket.com/checkout/checkout_payment/62?order_id=39982221&order_hash2a71d91259eefafd4bea3465050d4fa7ecb4186e',
                text: 'VA BRI',
                message: '',
                images: 'https://renan.tiket.com/images/apps_payment/bri_transfer.png',
                desc: 'Metode pembayaran ini menerima transfer melalui ATM, Internet Banking & Mobile Banking',
                countdown_time: 1061
              }
            ],
            countdown_time: 1061
          },
          {
            code: '59',
            link: 'https://renan.tiket.com/checkout/checkout_payment/59?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'ATM',
            message: '',
            type: 'atm_nicepay',
            desc: 'ATM Bersama, Prima & Alto',
            app_images: [
              'https://renan.tiket.com/images/apps_payment/atm_bersama.png',
              'https://renan.tiket.com/images/apps_payment/atm_prima.png',
              'https://renan.tiket.com/images/apps_payment/alto.png'
            ],
            countdown_time: 1061
          },
          {
            code: '3',
            link: 'https://renan.tiket.com/checkout/checkout_payment/3?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'KlikBCA',
            message: '',
            type: 'klikbca',
            desc: 'KlikBCA',
            app_images: ['https://renan.tiket.com/images/apps_payment/klik_bca.png'],
            countdown_time: 1061
          },
          {
            code: '4',
            link: 'https://renan.tiket.com/checkout/checkout_payment/4?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'BCA KlikPay',
            message: '',
            type: 'klikpay',
            desc: 'BCA Klikpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/bca_klikpay.png'],
            countdown_time: 1061
          },
          {
            code: '34',
            link: 'https://renan.tiket.com/checkout/checkout_payment/34?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'Mandiri Clickpay',
            message: 'Butuh No. Kartu dan Token',
            type: 'mandiri_clickpay',
            desc: 'Mandiri Clickpay',
            app_images: ['https://renan.tiket.com/images/apps_payment/mandiri_clickpay.png'],
            countdown_time: 1061
          },
          {
            code: '31',
            link: 'https://renan.tiket.com/checkout/checkout_payment/31?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'CIMB Clicks',
            message: '',
            type: 'cimbclicks',
            desc: 'CIMB Clicks',
            app_images: ['https://renan.tiket.com/images/apps_payment/cimb_click.png'],
            countdown_time: 1061
          },
          {
            code: '33',
            link: 'https://renan.tiket.com/checkout/checkout_payment/33?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            text: 'ePay BRI',
            message: '',
            type: 'epaybri',
            desc: 'ePay BRI',
            app_images: ['https://renan.tiket.com/images/apps_payment/e_pay_bri.png'],
            countdown_time: 1061
          }
        ],
        paymentDescription: {
          currency: 'IDR',
          orderId: '39982221',
          orderHash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
          grandTotal: 10001,
          subTotal: 1000,
          grandSubTotal: 3000,
          paymentCharge: 2000,
          baggageFee: [0],
          uniqueCode: 7001,
          orderExpiredDatetime: '2018-04-04 11:44:43',
          paymentType: 3,
          tixPoint: 0,
          giftPromo: false
        },
        loading: false,
        loaded: true
      },
      paymentDetail: {result: {}, loading: false, loaded: false},
      paymentGroup: {result: {}, loading: false, loaded: false},
      paymentConfirm: {
        data: {
          output_type: 'json',
          diagnostic: {
            status: 200,
            elapsetime: '0.7447',
            memoryusage: '38.13MB',
            unix_timestamp: 1522816063,
            confirm: 'success',
            lang: 'id',
            currency: 'IDR'
          },
          orderId: '39982221',
          result: {
            payment_subsider_tiket: 0,
            currency_to_be_converted: 'IDR',
            from_another_currency: false,
            reseller_id: null,
            payment_discount: 0,
            order_types: ['event'],
            order_id: '39982221',
            order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            currency: 'IDR',
            payment_charge: 2000,
            giftPromo: false,
            sub_total: 1000,
            unique_code: 7001,
            grand_total: 10001,
            grand_subtotal: 3000,
            orders: [
              {
                quantity: '1',
                order_detail_id: '58944361',
                order_type: 'event',
                event_name: 'Testing Event',
                tiket_name: 'Test CC',
                currency: 'IDR',
                price: 1000,
                tiket_event_start: '2012-02-01 00:00:00',
                tiket_event_end: '2018-06-30 00:00:00',
                event_type: 'B'
              }
            ],
            confirm_page_mobile: false,
            gaq: '',
            payment_type: 12,
            is_confirmation: true,
            is_change_payment: false,
            type: '',
            checkout_url: 'https://renan.tiket.com/checkout/checkout_payment/12?order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            arrOrderType: ['event'],
            order_expire_datetime: '2018-04-04 11:44:43',
            tiket_point: 0,
            tiket_point_worth: 0,
            tiket_point_notes: '',
            tiket_point_status: '',
            tiket_point_words: '',
            already_transfer_url: 'https://renan.tiket.com/confirmpayment/index/39982221/d8337c4e8ed29bf7494a0a37deebb39d'
          },
          banks: {
            bank_image: '/images/ico_bca.png',
            bank_owner_label: 'Nama',
            bank_owner: 'PT. Global Tiket Network',
            bank_name_label: 'Bank',
            bank_name: 'BCA',
            bank_branch_label: 'Cabang',
            bank_branch: 'Jakarta',
            bank_account_label: 'No Rekening',
            bank_account: '52 6032 2488'
          },
          message: 'Silakan transfer pembayaran untuk memesan. Anda memiliki waktu <strong> 17 menit  </strong> untuk melakukan pembayaran',
          grand_total: 10001,
          login_status: 'false',
          token: 'a5d579bfd821fa4bbba9c88a5fdf3f320b2affb7'
        },
        loading: false,
        loaded: true
      },
      paymentVerify: {data: {}, loading: false, loaded: false}
    },
    order: {
      result: {
        order_id: '39982221',
        data: [
          {
            expire: 19,
            order_detail_id: '58944361',
            order_expire_datetime: '2018-04-04 11:44:43',
            order_type: 'event',
            order_name: 'Testing Event',
            order_name_detail: 'Test CC',
            order_detail_status: 'active',
            detail: {
              order_detail_id: '58944361',
              tiket_barcode: '6546216472445',
              qty: '1',
              sell_rate_price: '1000.00',
              sell_rate_currency: 'IDR',
              tiket_seating: '',
              startdate: '2012-02-01 00:00:00',
              enddate: '2018-06-30 00:00:00',
              event_type: 'B',
              event_category: 'event',
              event_address: 'Lapangan D Senayan',
              travellers: [
                {
                  full_name: 'Abdul Rahman',
                  id_card: '09110110101',
                  salutation: 'Tuan',
                  seat: null
                }
              ],
              attend_date: null
            },
            order_photo: 'https://renan.tiket.com/img/business/f/u/business-future2.s.jpg',
            order_icon: 'h3d',
            tax_and_charge: '5000.00',
            subtotal_and_charge: '6000.00',
            delete_uri: 'https://renan.tiket.com/order/delete_order?order_detail_id=58944361&order_id=39982221&order_hash=2a71d91259eefafd4bea3465050d4fa7ecb4186e',
            business_id: '14161'
          }
        ],
        total_IDR: 1000,
        total: 6000,
        total_tax: 5000,
        total_without_tax: 1000,
        count_installment: 0,
        breakdown_price: [{title: '1 x Test CC', value: 1000}],
        promo: [],
        discount: 0,
        discount_amount: 0,
        total_without_promocode: 1000,
        payment_status: 'shoppingcart'
      },
      loading: false,
      loaded: true
    },
    app: {
      flash: {show: false, type: 'success', text: 'Total Payment Copied.'},
      popup: {show: false, header: '', footer: '', content: ''},
      account: {data: {}, loading: false, loaded: true},
      context: {
        query: {order_id: '39982221', order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'},
        params: {}
      }
    }
  });
});
