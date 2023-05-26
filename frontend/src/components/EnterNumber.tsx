import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {Button, Col, Form, Row} from 'react-bootstrap';
import { checkAnswers } from '../network/checkExtrasenseAnswers';

export default function EnterNumber() {
  const [number, setNumber] = useState<number>();
  const navigate = useNavigate();
  const handler = () => {
    if(number === undefined) return;
    checkAnswers(number);
    navigate("/");
  }
  return (
    <Row className='mt-4'>
      <Col></Col>
      <Col xs={4}>
        <Form>
          <Form.Group className="mb-3" controlId="formNumber">
            <Form.Label>Email address</Form.Label>
            <Form.Control type="input" placeholder="Enter guessed number" onChange={(e) => {setNumber(Number(e.target.value))}} />
            <Form.Text className="text-muted">
              It's just your guessed number, c'mon!
            </Form.Text>
          </Form.Group>
          <Button variant="primary" onClick={handler}>
            Send your number to server...
          </Button>
        </Form>
      </Col>
      <Col></Col>
    </Row>
  )
}