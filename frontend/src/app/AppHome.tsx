import { Box, Button, Card, Container, Flex, Grid, Group, Text, Title } from "@mantine/core";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { CommurzServiceClient } from "../service";
import { useDebouncedValue } from "@mantine/hooks";
import { useState } from "react";
import * as pb from '../pb/commurz/v1/commurz_pb'
import { ResultFromPromise } from "../model";
import { useNavigate } from "react-router-dom";

export function AppHome() {
	const defaultPageSelection = '10'
	const [name, setName] = useState('')
	const [debouncedName] = useDebouncedValue(name, 256);
	const [size, setSize] = useState<number | ''>(10);
	const [sizeSelect, setSizeSelect] = useState<string | null>(defaultPageSelection);
	const [page, setPage] = useState(1);
	const [total, setTotal] = useState(0)

	const navigate = useNavigate()

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

	const { data: resUser } = useQuery({
		queryKey: ['app', 'user'],
		queryFn: async () => {
			const res = await CommurzServiceClient.findUserByToken({})
			return res
		}
	})

	const queryClient = useQueryClient()

	const { data: resCart } = useQuery({
		queryKey: ['app', 'user-cart'],
		queryFn: async () => {
			const res = await CommurzServiceClient.findCartByUserToken({})
			return res
		}
	})

	const mutateAddToCart = useMutation({
		mutationKey: ['app', 'user-cart'],
		mutationFn: async (val: pb.AddProductToCartRequest) => CommurzServiceClient.addProductToCart(val),
		onSuccess: async () => {
			await queryClient.invalidateQueries(['app', 'user-cart'])
		}
	})

	if (isProductLoading) {
		return <Text>Loading...</Text>
	}

	if (error) {
		return alert(error)
	}

	function getItemInCart() {
		return resCart?.items?.length ?? 0
	}

	async function buyProduct(productID: string) {
		const req = new pb.AddProductToCartRequest({
			productId: productID,
			quantity: BigInt(1),
			userId: resUser?.id ?? '',
		})
		const res = await ResultFromPromise(mutateAddToCart.mutateAsync(req))
		if (!res.ok) {
			alert(res.error.message)
			return
		}
		alert('Product added to cart')
	}

	function isProductOOS(product: pb.Product) {
		return product.currentStock <= 0
	}

	return <>
		<Container>
			<Flex align="center" justify="space-between">
				<Title>Products</Title>
				<Group pr="sm">
					<Text size="xl">Cart: {getItemInCart()}</Text>
					<Button size="sm" onClick={() => {
						navigate('/checkout')
					}}>Checkout</Button>
				</Group>
			</Flex>

			<Grid py="md" grow>
				{resListProducts?.products.map((product) => {
					return <>
						<Grid.Col span={4}>
							<Card key={product.id} padding="md" radius="md" withBorder maw={300}
								opacity={isProductOOS(product) ? 0.7 : 1}>
								<Card.Section>
									<img src="https://picsum.photos/300" alt={product.name} />
								</Card.Section>
								<Card.Section p="md">
									<Text size="xl" pb="sm">{product.name}</Text>
									<Group spacing="xl" align="flex-start">
										<Text size="md">IDR {product.textPriceIdr}</Text>
										<Text size="md">Stock: {product.currentStock.toString()}</Text>
									</Group>
								</Card.Section>
								<Card.Section p="md">
									<Button variant="light" fullWidth color="green.9"
										disabled={isProductOOS(product)}
										onClick={() => buyProduct(product.id)}>Buy</Button>
								</Card.Section>
							</Card>
						</Grid.Col>
					</>
				})}
			</Grid>
		</Container>
	</>;
}