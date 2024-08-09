import {
	Navbar as BootstrapNavbar,
	Container,
	Nav,
	NavDropdown,
} from 'solid-bootstrap'

export const Navbar = () => {
	return (
		<BootstrapNavbar bg="body-tertiary">
			<Container fluid>
				<BootstrapNavbar.Brand href="#">
					<i class="bi bi-house"></i>
					{' Home Manager'}
				</BootstrapNavbar.Brand>
				<BootstrapNavbar.Toggle aria-controls="navbar" />
				<BootstrapNavbar.Collapse id="navbar">
					<Nav class="me-auto">
						<Nav.Link href="">Rooms</Nav.Link>
						<Nav.Link href="">Assignments</Nav.Link>
					</Nav>
				</BootstrapNavbar.Collapse>
			</Container>
		</BootstrapNavbar>
	)
}
