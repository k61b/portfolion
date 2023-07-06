import "@styles/globals.css";
import { i18n } from "@utils/i18n";

export const metadata = {
  title: "Portfolion",
  description: "Your portfolio manager",
};

export async function generateStaticParams() {
  return i18n.locales.map((locale) => ({ lang: locale }))
}

export default function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { lang: string };
}) {
  return (
    <html lang={params.lang}>
      <body>{children}</body>
    </html>
  );
}
