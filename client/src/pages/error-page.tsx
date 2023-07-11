import { useRouteError } from "react-router-dom";
import { IError } from "@/types/error";
import { Link } from "react-router-dom";

export default function ErrorPage() {
  const error = useRouteError() as IError;
  console.log(error);

  return (
    <div
      id="error-page"
      className="flex flex-col justify-center items-center h-screen"
    >
      <h1>Oops!</h1>
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>{error.statusText || error.message}</i>
        <Link
          to={"/"}
          className="italic underline underline-offset-1 text-purple-800"
        >
          Go to the home page
        </Link>
      </p>
    </div>
  );
}
