import React from 'react';
import { Container, Nav, Navbar } from 'react-bootstrap';

export const Navigation = () => {
  return (
    <Navbar bg="light" expand="lg">
      <Container>
        <Navbar.Brand href="/">ExtraSense</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="/history">Extrasenses guesses history</Nav.Link>
            <Nav.Link href="/trust">Extrasenses trust</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};