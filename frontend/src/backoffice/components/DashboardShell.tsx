import { AppShell, Header, NavLink, Navbar, Text } from "@mantine/core";
import { PropsWithChildren } from "react";
import { useNavigate } from "react-router-dom";

type Props = PropsWithChildren & {
    activeTab?: '' | 'products'
}

export function DashboardShell(props: Props) {
    const navigate = useNavigate();

    return (
        <AppShell
        padding="md"
        header={
            <Header height={60} p="xs">
                <Text size="lg">Backoffice Products</Text>
            </Header>
        }
        navbar={
            <Navbar width={{ base: 300 }} height={500} p="xs">
                <NavLink label="Products" active={props.activeTab === 'products'} onClick={() => {
                    navigate('/backoffice/products')
                }}>
                </NavLink>
            </Navbar>
        }
        styles={(theme) => ({
            main: { backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0] },
        })}
        >
        {props.children}
    </AppShell>
    )
}