import { Locale, getDictionary } from '@utils/i18n'

export default async function Dashboard({
  params: { lang },
}: {
  params: { lang: Locale }
}) {
  const dictionary = await getDictionary(lang)
  return (
    <div>
      <h1>{dictionary['server-component'].welcome}</h1>
    </div>
  )
}
