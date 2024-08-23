import { Button, Container, Form, Row } from 'solid-bootstrap'
import { Role } from '../../models/Role'
import { createSignal } from 'solid-js'
import { NewUserRequest } from '../../models/Users'
import UsersRepository from '../../repositories/usersRepository'

export interface CreateUserProps {
	userRepository: UsersRepository
}

export const CreateUser = (props: CreateUserProps) => {
	const [validated, setValidated] = createSignal(false)
	const [name, setName] = createSignal('')
	const [role, setRole] = createSignal(Role.NotSet)
	const userRepository = props.userRepository

	const createUser = async () => {
		const newUserRequestResult = NewUserRequest.from(name(), role())

		if (newUserRequestResult.isErr()) {
			console.log('Validation errors while making newUserRequest', newUserRequestResult.err?.toString())
			return
		}

		const result = await userRepository.createUser(newUserRequestResult.ok!.val)

		if (result.isErr()) {
			console.log(result.err?.message)
		}

		if (result.isOk()) {
			console.log(JSON.stringify(result.ok!.val))
		}
	}

	const handleSubmit = (event: SubmitEvent) => {
		const form = event.currentTarget
		event.preventDefault()

		setValidated(true)
		if ((form as HTMLFormElement).checkValidity() === false) {
			event.stopPropagation()
		} else {
			createUser()
		}
	}

	return (
		<>
			<Container>
				<Form
					noValidate
					validated={validated()}
					onSubmit={handleSubmit}
				>
					<Row class="mb-3">
						<Form.Group class="mb-3" controlId="formUser">
							<Form.Label>Name</Form.Label>
							<Form.Control
								type="text"
								placeholder="Name"
								value={name()}
								onChange={(event) =>
									setName(event.target.value)
								}
								required
							/>
						</Form.Group>
						<Form.Group>
							<Form.Label>Role</Form.Label>
							<Form.Select
								value={role()}
								onChange={(event) =>
									setRole(event.target.value as Role)
								}
								required
							>
								<option value="" disabled>
									Select role
								</option>
								<option value={Role.Admin}>{Role.Admin}</option>
								<option value={Role.User}>{Role.User}</option>
							</Form.Select>
						</Form.Group>
					</Row>
					<Button variant="primary" type="submit">
						Submit
					</Button>
				</Form>
			</Container>
		</>
	)
}
