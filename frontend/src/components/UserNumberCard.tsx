import React from 'react';
import { Card } from 'react-bootstrap';
import { TransactionHistory } from '../types/dto/transaction_history';


export const UserNumberCard = (props: TransactionHistory) => {
  return (
    <Card>
      <Card.Body>
        <Card.Title>Number {props.correct_number}</Card.Title>
        <Card.Text>
          {props.extrasense_guesses.map((value) => {
            return `${value.guessed_by}: ${value.number} `
          })}
        </Card.Text>
      </Card.Body>
    </Card>
  )
}