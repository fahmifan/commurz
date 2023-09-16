import { Button, Card, Container, Grid, Text, Title } from "@mantine/core";
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

	if (isProductLoading) {
		return <Text>Loading...</Text>
	}

	if (error) {
		return alert(error)
	}

  return <>
	<Container>
		<Title>Products</Title>
		
		<Grid py="md">
		{resListProducts?.products.map((product) => {
			return <>
				<Grid.Col span={4}>
					<Card key={product.id} padding="lg" radius="md" withBorder maw={300}>
						<Card.Section>
							<img src="https://picsum.photos/300" alt={product.name} />
						</Card.Section>
						<Card.Section p="md">
							<Text size="md">{product.name}</Text>
							<Text size="md">IDR {product.textPriceIdr}</Text>
						</Card.Section>
						<Card.Section p="md">
							<Button variant="light" fullWidth color="green.9">Buy</Button>
						</Card.Section>
					</Card>
				</Grid.Col>
			</>
		})}
		</Grid>
	</Container>
  </>;
}