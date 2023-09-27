import { 
    createPromiseClient, 
    Interceptor,
} from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

import { CommurzService } from "../pb/commurz/v1/commurz_connect"
import { Result, ResultFromPromise, ResultFromResponse, OK, Err } from "../model";

const host = ""

export const CommurzServiceClient = createPromiseClient(CommurzService, transport());

function transport() {
    return createConnectTransport({
        baseUrl: "/grpc",
        interceptors: [csrfInterceptor()],
    })
}

function csrfInterceptor(): Interceptor {
    return (next) => async (req) => {
        const csrfRes = await getCSRFToken();
        if (!csrfRes.ok) {
            throw csrfRes.error;
        }

        req.header.set("X-Csrf-Token", csrfRes.value.csrf);
        return await next(req);
    };
}

type GetCSRFResponse = {
    csrf: string
}

export async function getCSRFToken(): Promise<Result<GetCSRFResponse, Error>> {
    const res = await ResultFromPromise(fetchJSON('GET', `${host}/api/auth/csrf`))
    if (!res.ok) {
        return Err(res.error)
    }

    const csrfRes = await ResultFromResponse<GetCSRFResponse>(res.value)
    if (!csrfRes.ok) {
        return Err(csrfRes.error)
    }

    return OK(csrfRes.value)
}

function fetchJSON<T>(method: string, url: string, body?: T) {
    return fetch(url, {
        method,
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
    })
}
