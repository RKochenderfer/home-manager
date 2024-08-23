import { HttpMethod } from '../models/HttpMethod'
import { NewUserRequest, NewUserResponse } from '../models/Users'
import { Result } from '../result'
import { Repository } from './repository'

export default class UsersRepository extends Repository {
	private constructor() {
		super('/users')
	}

	async createUser(
		newUser: NewUserRequest,
	): Promise<Result<NewUserResponse>> {
		try {
			const result = await fetch(super.endpoint, {
				method: HttpMethod.POST,
				body: JSON.stringify(newUser),
				headers: super.getHeaders(),
			})

			if (!result.ok) {
				return Result.err(
					`Failed to create user with http status code: ${result.status} and text: ${result.statusText}`,
				)
			}

			const newUserResponse: NewUserResponse = await result.json()

			return Result.ok(newUserResponse)
		} catch (error) {
			return Result.err('Failed to create user', error as Error)
		}
	}
}
