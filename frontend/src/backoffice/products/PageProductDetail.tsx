import { Anchor, Box, Breadcrumbs, Button, Group, Loader, Modal, NumberInput, Pagination, Select, Table, Text, TextInput, Title } from "@mantine/core";

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { CommurzServiceClient } from "../../service";
import { Form, useForm } from "@mantine/form";

import * as pb from '../../pb/commurz/v1/commurz_pb'
import { DashboardShell } from "../components/DashboardShell";
import { ResultFromPromise } from "../../model";

export function PageProductDetail() {
    const searchParam = new URLSearchParams(window.location.search)
    const productID = searchParam.get('product_id')
    if (!productID) {
        return <DashboardShell>
            <Title>Product not found</Title>
        </DashboardShell>
    }

    const formUpdateProductStock = useForm({
        initialValues: {
            stockIn: 0,
            stockOut: 0,
        },
    })

    const queryClient = useQueryClient()
    const { isLoading: isFindProductLoading, error: findProductErr, data: resFindProduct } = useQuery({
        queryKey: ['backoffice', 'products', 'id', productID],
        queryFn: async () => {
            const res = await CommurzServiceClient.findProductByID({
                id: productID ?? '',
            })

            return res
        },
    })
    const mutationUpdateStock = useMutation({
        mutationKey: ['backoffice', 'products', 'id', productID],
        mutationFn: async (val: pb.UpdateProductStockRequest) => CommurzServiceClient.updateProductStock(val),
    })

    if (isFindProductLoading) {
        return <DashboardShell>
            <Title>Loading...</Title>
        </DashboardShell>
    }

    if (findProductErr) {
        return <DashboardShell>
            <Title>Product not found</Title>
        </DashboardShell>
    }

    async function onSubmitUpdateStock(values: {
        stockIn: number,
        stockOut: number,
    }) {
        const req = new pb.UpdateProductStockRequest({
            productId: productID ?? '',
            stockIn: BigInt(values.stockIn),
            stockOut: BigInt(values.stockOut),
            version: resFindProduct?.version ?? BigInt(0),
        })
        const res = await ResultFromPromise(mutationUpdateStock.mutateAsync(req))
        if (!res.ok) {
            alert(res.error.message)
            return
        }

        await queryClient.invalidateQueries(['backoffice', 'products', 'id', productID])
        formUpdateProductStock.reset()

        alert('Success add stock')
    }

    return <DashboardShell activeTab="">
        <Breadcrumbs>
            <Anchor href="/backoffice/products" key={1}>
                Products
            </Anchor>
            <Anchor href="#" key={1}>
            {resFindProduct?.id}
            </Anchor>
        </Breadcrumbs>

        <br />
        <br />

        <Text size="lg" weight="bold">Product Detail</Text>
        <Box maw="300px">
            <Text pt="sm" pb="sm">Name: {resFindProduct?.name}</Text>
        </Box>

        <br />
        <br />

        <Text size="lg" weight="bold">Update Stock</Text>
        <Box maw="300px">
            <Form form={formUpdateProductStock} onSubmit={onSubmitUpdateStock}>
                <Text pt="sm" pb="sm">Current Stock: {resFindProduct?.currentStock.toLocaleString()}</Text>
                <NumberInput label="Stock In" placeholder="Stock"
                    {...formUpdateProductStock.getInputProps('stockIn')} />
                <NumberInput label="Stock Out" placeholder="Stock"
                    {...formUpdateProductStock.getInputProps('stockOut')} />

                <br />
                <Button type="submit">Submit</Button>
            </Form>

        </Box>
    </DashboardShell>
}
