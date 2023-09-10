import { Text } from "@mantine/core";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { CommurzServiceClient } from "../service";
import { useDebouncedValue } from "@mantine/hooks";
import { useState } from "react";

export function AppHome() {
	const defaultPageSelection = '10'
	const pageSelections = ['5', '10', '20', '50']
	const [name, setName] = useState('')
	const [debouncedName] = useDebouncedValue(name, 256);
	const [size, setSize] = useState<number | ''>(10);
	const [sizeSelect, setSizeSelect] = useState<string | null>(defaultPageSelection);
	const [page, setPage] = useState(1);
	const [total, setTotal] = useState(0)  
	
  const { isLoading: isProductLoading, error, data: resListProducts } = useQuery({
		queryKey: ['backoffice', 'products', page, size, debouncedName],
		keepPreviousData: true,
		queryFn: async () => {
			const res = await CommurzServiceClient.listBackofficeProducts({
				name: debouncedName,
				pagination: {
					page: page,
					size: size as number,
				},
			})

			const pageTotal = Math.ceil(res.count / (size as number))
			setTotal(pageTotal)

			return res
		},
	})

  return <Text>Home</Text>;
}