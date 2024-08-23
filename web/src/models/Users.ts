import { Result } from '../result'
import { Role } from './Role'

export class NewUserRequest {
	private constructor(
		private readonly name: string,
		private readonly role: string,
	) {}

	public static from(name: string, role: Role): Result<NewUserRequest> {
		if (name.trim().length === 0) {
			return Result.err('Name cannot be empty')
		}

		if (role === Role.NotSet) {
			return Result.err('Role cannot be NotSet')
		}

		return Result.ok(new NewUserRequest(name, role))
	}
}

export interface NewUserResponse {}
