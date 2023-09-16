import { Button, Group, Loader, Modal, NumberInput, Pagination, Select, Table, TextInput } from "@mantine/core";

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { CommurzServiceClient } from "../../service";
import { FormEvent, useState } from "react";
import { useDebouncedValue, useDisclosure } from "@mantine/hooks";
import { useForm } from "@mantine/form";

import * as pb from '../../pb/commurz/v1/commurz_pb'
import { DashboardShell } from "../components/DashboardShell";
import { Link, redirect } from "react-router-dom";

export default function PageProducts() {
	const defaultPageSelection = '10'
	const pageSelections = ['5', '10', '20', '50']
	const [name, setName] = useState('')
	const [debouncedName] = useDebouncedValue(name, 256);
	const [size, setSize] = useState<number | ''>(10);
	const [sizeSelect, setSizeSelect] = useState<string | null>(defaultPageSelection);
	const [page, setPage] = useState(1);
	const [total, setTotal] = useState(0)

	const formAddProduct = useForm({
		initialValues: {
			name: '',
			price: 0,
		},
	})
	const [addProductOpened, { open: openAddProductModal, close: closeAddproductModal }] = useDisclosure();

	const queryClient = useQueryClient()
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
	const mutationCreateProduct = useMutation({
		mutationKey: ['backoffice', 'products'],
		mutationFn: async (val: pb.CreateProductRequest) => CommurzServiceClient.createProduct(val),
		onSuccess: async () => {
			await invalidateProductQuery()
		},
	})

	async function invalidateProductQuery() {
		await queryClient.invalidateQueries(['backoffice', 'products'])
	}

	async function onSubmitProduct(e: FormEvent<HTMLFormElement>) {
		e.preventDefault()

		const { name, price } = formAddProduct.values
		await mutationCreateProduct.mutateAsync(new pb.CreateProductRequest({
			name,
			price: BigInt(price),
		}))

		formAddProduct.reset()
		closeAddproductModal()
	}

	if (error) {
		return <>
			{() => {
				alert('Something went wrong!')
			}}
		</>
	}

	if (isProductLoading) {
		return <Loader />
	}

	return <DashboardShell activeTab="products">
		<Group>
			<Button onClick={openAddProductModal}>Add Product</Button>
			<TextInput placeholder="Search..." onChange={(ev) => {
				setName(ev.target.value)
			}} />
		</Group>

		<Table verticalSpacing="sm">
			<thead>
				<tr>
					<th>Name</th>
					<th>Price</th>
					<th>Current Stock</th>
					<th>Actions</th>
				</tr>
			</thead>
			<tbody>{resListProducts?.products.map(prod => (
				<tr key={prod.id}>
					<td>{prod.name}</td>
					<td>{prod.textPriceIdr}</td>
					<td>{prod.currentStock.toString()}</td>
					<td>
						<Link	to={`/backoffice/products/stocks?product_id=${prod.id}`}>
							Update Stock
						</Link>
					</td>
				</tr>
			))}</tbody>
		</Table>

		<Group>
			<Pagination total={total} onChange={setPage} />
			<Select label=""
				maxLength={3}
				data={pageSelections}
				value={sizeSelect}
				onChange={(val: string) => {
					setSizeSelect(val)
					setSize(parseInt(val))
					setPage(1)
				}}
			/>
		</Group>

		{/* Modals */}
		<Modal opened={addProductOpened} onClose={closeAddproductModal}
			title="Add Product">
			<form onSubmit={onSubmitProduct}>
				<TextInput
					placeholder="Name"
					pb="sm"
					label="Name"
					{...formAddProduct.getInputProps('name')}
				/>

				<NumberInput placeholder="Price"
					label="Price"
					pb="sm" 
					{...formAddProduct.getInputProps('price')}
				/>

				<Button type="submit">Add</Button>
			</form>
		</Modal>
	</DashboardShell>
}
