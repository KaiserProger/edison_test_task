import React from 'react';
import { useNavigate } from 'react-router-dom';
import {Button, Col, Form, Row} from 'react-bootstrap';
import { checkAnswers } from '../network/checkExtrasenseAnswers';
import { Formik } from 'formik';
import * as Yup from 'yup';

const ERROR_MESSAGE = 'Please enter positive two-digit number!';

const numberSchema = Yup.object().shape({
  number: Yup.number()
    .min(10, ERROR_MESSAGE)
    .max(99, ERROR_MESSAGE)
    .required(),
});

export default function EnterNumber() {
  const navigate = useNavigate();
  return (
    <Row className='mt-4'>
      <Col></Col>
      <Col xs={4}>
        <h1>Enter two-digit number</h1>
        <Formik
          initialValues={{number: 10}}
          validationSchema={numberSchema}
          onSubmit={values => {
            checkAnswers(values.number);
            navigate("/");
          }}
        >
          {({handleSubmit, handleChange, values, errors, touched}) => (
            <Form noValidate onSubmit={handleSubmit}>
              <Form.Group className="mb-3" controlId="formNumber">
                <Form.Label>A number</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Enter guessed number"
                  name="number"
                  onChange={handleChange} 
                  value={values.number}
                  isInvalid={!!errors.number}
                />
                <Form.Control.Feedback type="invalid">
                  {errors.number}
                </Form.Control.Feedback>
              </Form.Group>
              <Button variant="primary" type='submit'>
                Send your number to server...
              </Button>
            </Form>
          )}
        </Formik>
      </Col>
      <Col></Col>
    </Row>
  );
};
