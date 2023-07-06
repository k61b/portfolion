import AuthForm from "@app/[lang]/(auth)/components/auth";
import { Locale } from "@utils/i18n";

export default async function Login({
  params: { lang },
}: {
  params: { lang: Locale };
}) {

  return (
    <div className="flex flex-col justify-center items-center  h-screen max-h-screen dark:bg-slate-800">
      <AuthForm lang={lang} />
    </div>
  );
}
