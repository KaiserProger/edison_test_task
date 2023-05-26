import React from 'react';
import { Button } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';
import { init } from '../network/init';

export default function StartButton() {
  const navigate = useNavigate();
  const handler = async () => {
    await init();
    navigate("/poll");
  }
  return (
    <>
      <Button onClick={handler}>
        Start your way!
      </Button>
    </>
  )
}