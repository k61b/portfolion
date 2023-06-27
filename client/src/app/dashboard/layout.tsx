import Navbar from "@/components/navbar";

export const metadata = {
  title: "Dashboard",
  description: "Dashboard for bookmarks",
};

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <Navbar />
        {children}
      </body>
    </html>
  );
}
