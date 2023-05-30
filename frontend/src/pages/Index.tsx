import React from 'react';
import { Col, Container, Row } from 'react-bootstrap';
import StartButton from '../components/StartButton';
import { Navigation } from '../components/Navigation';

export const Index = () => {
  return (
    <>
      <Navigation/>
      <Container className="fluid h-75">
          <Row className='h-100'>
            <Col className='d-flex flex-column align-self-center justify-content-center'>
              <h3>Press button</h3>
              <StartButton></StartButton>
            </Col>
          </Row>
      </Container>
      </>
  );
};