import "./globals.css";

export const metadata = {
  title: "Portfolion",
  description: "Your portfolio manager",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
