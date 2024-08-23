export abstract class Repository {
	constructor(protected endpoint: string) {}

	protected getHeaders(): Headers {
		const headers = new Headers()

		headers.set('Content-Type', 'application/json')

		return headers
	}
}
