import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Box, Button, Card, Divider, Stack, TextField } from '@mui/material';
import { IProduct } from './Products';
import { useNavigate, useParams } from 'react-router-dom';

const initValue = {
  id: 0,
  title: '',
  description: '',
  image: '',
  price: 0,
};

export default function EditProduct() {
  const navigate = useNavigate();
  const { productId } = useParams();
  const [productData, setProductData] = useState<IProduct>(initValue);

  const handleClick = () => {
    axios.put(`/product/${productId}`, productData).then(() => navigate('/products'));
  };

  useEffect(() => {
    axios.get(`/product/${productId}`).then(({ data }) => setProductData(data));
  }, []);

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
            <TextField value={productData.title} name="title" label="Title" onChange={onChange} />
            <TextField value={productData.description} name="description" multiline rows="4" label="Description" onChange={onChange} />
            <TextField value={productData.image} name="image" label="Image" onChange={onChange} />
            <TextField value={productData.price}  name="price" label="Price" type="number" onChange={onChange} />
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
