import i18n, { getTranslation, sanitizeLang } from '../lang';

const langTest = 'Test';

test('Should function localization correctly', () => {
  let test = i18n(langTest);

  expect(test).toEqual(langTest);
  test = i18n('');
  expect(test).toEqual('');
  test = i18n('test.withParams', 'lalala');
  expect(test).toEqual('test lalala');
  test = i18n('test.withParams');
  expect(test).toEqual('test {0}');
  test = getTranslation(undefined, 'test');
  expect(test).toEqual('test');
  let sanitizedLang = sanitizeLang('id');

  expect(sanitizedLang).toEqual('id');
});
