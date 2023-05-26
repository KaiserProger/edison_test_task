import React from 'react'
import { GuessTransaction } from '../types/dto/guess_transaction'
import { Card } from 'react-bootstrap'


export const ExtraSenseCard = (props: GuessTransaction) => {
  return (
    <Card style={{ width: '18rem', }} className="mx-2">
      {/* <Card.Img variant="top" src="holder.js/100px180" />*/}
      <Card.Body>
        <Card.Title>Extrasense {props.guessed_by}</Card.Title>
        <Card.Text>
          He thinks your number is {props.number}
        </Card.Text>
      </Card.Body>
    </Card>
  )
}