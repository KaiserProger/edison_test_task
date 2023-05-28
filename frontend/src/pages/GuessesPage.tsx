import React, { useEffect, useState } from 'react'
import { pollExtrasenses } from '../network/pollExtrasenses';
import { GuessTransaction } from '../types/dto/guess_transaction';
import { ExtraSenseCard } from '../components/ExtraSenseCard';
import { Col, Container, Row } from 'react-bootstrap';
import EnterNumber from '../components/EnterNumber';
import { Navigation } from '../components/Navigation';

export const GuessesPage = () => {
  const [transactions, setTransactions] = useState<GuessTransaction[]>([]);
  useEffect(() => {
    const handler = async () => {
      await pollExtrasenses();
      setTransactions(JSON.parse(sessionStorage.getItem("guesses") || "[]"));
    };
    handler();
  }, [setTransactions]);
  if(transactions.length === 0) return (
    <>
      <p>Please wait!</p>
    </>
  )
  return (
    <>
      <Navigation/>
      <Container>
        <Row>
          <Col className='d-flex flex-row'>
            {transactions.map((value, index) => {
              return <ExtraSenseCard {...value} key={index}></ExtraSenseCard>
            })}
          </Col>
        </Row>
        <EnterNumber></EnterNumber>
      </Container>
    </>
  );
};
