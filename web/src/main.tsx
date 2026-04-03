import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ConfirmationPage } from "./ConfirmationPage.tsx";
import { SinglePost } from "./SinglePost.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/feed",
    element: <App />,
  },
  {
    path: "/post/:postID",
    element: <SinglePost />,
  },
  {
    path: "/confirm/:token",
    element: <ConfirmationPage />,
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
);
