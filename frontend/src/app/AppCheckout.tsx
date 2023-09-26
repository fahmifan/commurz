import { Box, Button, Card, Container, Group, Text, Title } from "@mantine/core";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import React from "react";
import { CommurzServiceClient } from "../service";
import { useNavigate } from "react-router-dom";
import { ResultFromPromise } from "../model";

export function AppCheckout(): React.ReactNode {
    const navigate = useNavigate()
    const queryClient = useQueryClient()

    const { data: resCart } = useQuery({
        queryKey: ['app', 'cart'],
        queryFn: async () => await CommurzServiceClient.findCartByUserToken({}),
    })

    const { data: resUser } = useQuery({
        queryKey: ['app', 'user'],
        queryFn: async () => await CommurzServiceClient.findUserByToken({}),
    })

    const mutateCheckoutAll = useMutation({
        mutationKey: ['app', 'cart'],
        mutationFn: async () => await CommurzServiceClient.checkoutAll({
            userId: resUser?.id as string,
        }),
        onSuccess: async () => {
            queryClient.invalidateQueries(['app', 'cart'])
            queryClient.invalidateQueries(['app', 'products'])    
        },
    })

    async function onOrder() {
        const res = await ResultFromPromise(mutateCheckoutAll.mutateAsync())
        if (!res.ok) {
            alert(res.error)
            return
        }

        alert('Order success')
        navigate('/')
    }

    return <>
    <Container>
        <Title>Checkout</Title>
        {resCart && resCart?.items.map((item, index) => {
            return <Card key={index}>
                <Group>
                    <img src="https://picsum.photos/50" alt={item.product?.name} />
                    <Box>
                        <Text>{item.product?.name}</Text>
                        <Text>{item.productPrice.toString()} x {item.quantity.toString()}</Text>
                    </Box>
                </Group>
            </Card>
        })}
        <Button onClick={() => { onOrder() }}>Order</Button>
    </Container>
    </>
}