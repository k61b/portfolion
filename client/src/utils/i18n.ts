export const i18n = {
  defaultLocale: 'en',
  locales: ['en', 'tr'],
} as const

export type Locale = (typeof i18n)['locales'][number]

const dictionaries = {
  en: () => import('../locales/en.json').then((module) => module.default),
  tr: () => import('../locales/tr.json').then((module) => module.default),
}

export const getDictionary = async (locale: Locale) => dictionaries[locale]()
