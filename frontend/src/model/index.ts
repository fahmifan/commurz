export type Result<T, E = string> = { value: T, ok: true } | { error: E, ok: false }


export async function ResultFromPromise<T>(prom: Promise<T>): Promise<Result<T, any>> {
    try {
        const data = await prom;
        return OK(data);
    } catch (error) {
        return Err(error);
    }
}

export async function ResultFromResponse<T>(res: Response): Promise<Result<T, any>> {
    const data = await ResultFromPromise(res.json())
    if (!data.ok) {
        return Err(data.error)
    }

    return OK(data.value as T)
}

export function Err<E>(why: E): Result<any, E> {
    return {
        ok: false,
        error: why
    }
}

export function OK<T>(value: T): Result<T, any> {
    return {
        ok: true,
        value: value
    }
}