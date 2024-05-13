import { createI18n } from 'vue-i18n'

const enUS = {
    hello: 'Hello, {name}!'
}

type MessageSchema = typeof enUS

const i18n = createI18n<[MessageSchema], 'en-US'>({
  locale: 'en-US',
  messages: {
    'en-US': enUS
  }
})

export default i18n;