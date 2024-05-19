import React, { useState } from 'react';
import axios from 'axios';
import { Box, Button, Card, Divider, Stack, TextField } from '@mui/material';
import { IProduct } from './Products';
import { useNavigate } from 'react-router-dom';

const initValue = {
  title: '',
  description: '',
  image: '',
  price: 0,
};

export default function CreateProduct() {
  const navigate = useNavigate();
  const [productData, setProductData] = useState<Omit<IProduct, 'id'>>(initValue);

  const handleClick = () => {
    axios.post('/product', productData).then(() => navigate('/products'));
  };

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    let { name, value } = e.target;
    let val: string | number = value;
    if (name === 'price') {
      val = Number(val);
    }
    setProductData((state) => ({
      ...state,
      [name]: val,
    }));
  };

  return (
    <Box sx={{ height: 1 }}>
      <Stack alignItems="center" justifyContent="center" sx={{ height: 1 }}>
        <Card
          sx={{
            p: 5,
            width: 1,
            maxWidth: 420,
          }}
        >
          <Stack spacing={3}>
            <TextField name="title" label="Title" onChange={onChange} />
            <TextField name="description" multiline rows="4" label="Description" onChange={onChange} />
            <TextField name="image" label="Image" onChange={onChange} />
            <TextField name="price" label="Price" type="number" onChange={onChange} />
          </Stack>

          <Divider sx={{ mt: 2, mb: 2 }} />

          <Button
            fullWidth
            size="large"
            type="submit"
            variant="contained"
            color="primary"
            onClick={handleClick}
          >
            Create Product
          </Button>
        </Card>
      </Stack>
    </Box>
  );
}
