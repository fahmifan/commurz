import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

import { MantineProvider, Text, Title } from "@mantine/core";
import BackofficeProducts from "./backoffice/BackofficeProducts";
import { AppHome } from "./app/AppHome";


const router = createBrowserRouter([
  {
    path: "/",
    element: <AppHome />
  },
  {
    path: "/backoffice/products",
    element: <BackofficeProducts />,
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