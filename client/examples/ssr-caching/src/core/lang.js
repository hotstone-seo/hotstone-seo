import id from '../lang/id.json';
import en from '../lang/en.json';

const lang = {
  id: id,
  en: en
};

if (!String.prototype.format) {
  String.prototype.format = function() {
    const args = arguments;

    return this.replace(/{(\d+)}/g, function(match, number) {
      return typeof args[number] !== 'undefined' ? args[number] : match;
    });
  };
}

export const sanitizeLang = lang => {
  return /^(id|en)$/.test(lang) ? lang : 'id';
};

export const getTranslation = (source, key) => {
  if (!key) {
    return '';
  }

  if (!source) {
    return key;
  }

  return source[key] || key;
};

const findTranslation = (source, key) => {
  if (!key) {
      return '';
  }

  const matches = key.split('.');
  let i = 0;
  let str = source;

  while (i < matches.length) {
    str = getTranslation(str, matches[i]);
    i++;
  }

  return str;
};

// Example usage
// don't forget to add it on the [lang].json
// {i18n('Halo {0} asdasdasdasdsadsa {1} !', 'oka lalala', 'okok')}
export default (cookieLang = 'id') => {
  const currentLang = lang[sanitizeLang(cookieLang)];

  return (str, ...rest) => {
    const translation = findTranslation(currentLang, str);

    return translation && typeof translation === 'string' ? translation.format(...rest) : str;
  };
};
