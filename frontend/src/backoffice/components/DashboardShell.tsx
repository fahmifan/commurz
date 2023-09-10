import { AppShell, Header, NavLink, Navbar, Text } from "@mantine/core";
import { PropsWithChildren } from "react";
import { Link } from "react-router-dom";

export function DashboardShell(props: PropsWithChildren) {
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
                <Link to="/backoffice/products" style={{textDecoration: 'none'}}>
                    <NavLink label="Products" />
                </Link>
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