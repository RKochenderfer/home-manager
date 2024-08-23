class Ok<T> {
	constructor(private readonly _val: T) {}

	get val(): T {
		return this._val
	}
}

class Err {
	constructor(
		private readonly _message: string,
		private readonly _err?: Error,
	) {}

	get message(): string {
		return this._message
	}

	get err(): Error | undefined {
		return this._err
	}

	toString(): string {
		if (this._err !== undefined) {
			return `message: ${this._message} | Error ${JSON.stringify(this._err)}`
		}
		return `message: ${this._message}`
	}
}

export class Result<T> {
	private readonly _ok: Ok<T> | undefined
	private readonly _err: Err | undefined

	private constructor(ok?: Ok<T>, err?: Err) {
		this._ok = ok
		this._err = err
	}

	isErr() {
		return this._err !== undefined
	}

	isOk() {
		return this._ok !== undefined
	}

	static ok<T>(val: T): Result<T> {
		const ok = new Ok(val)
		return new Result(ok)
	}

	static err<T>(msg: string, error?: Error): Result<T> {
		return new Result<T>(undefined, new Err(msg, error))
	}

	get ok(): Ok<T> | undefined {
		return this._ok
	}

	get err(): Err | undefined {
		return this._err
	}
}
