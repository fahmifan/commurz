import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

import { MantineProvider } from "@mantine/core";
import backoffices from "./backoffice";
import { AppHome } from "./app/AppHome";

const router = createBrowserRouter([
  {
    path: "/",
    element: <AppHome />
  },
  {
    path: "/backoffice/products",
    element: <backoffices.PageProducts />,
  },
  {
    path: "/backoffice/products/detail",
    element: <backoffices.PageProductDetail />,
  },
]);

function App() {
  const queryClient = new QueryClient()
  
  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider>
        <RouterProvider router={router} />
      </MantineProvider>
    </QueryClientProvider>
  )
}

export default App