import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Box, Button, ButtonGroup, Table, TableCell, TableContainer, TableFooter, TablePagination, TableRow } from '@mui/material';
import { useNavigate } from 'react-router-dom';

export interface IProduct {
  id: number;
  title: string;
  description: string;
  image: string;
  price: number;
}

export default function Products() {
  const navigate = useNavigate();
  const [products, setProducts] = useState<IProduct[]>([]);
  const [page, setPage] = useState(0);
  const [perPage, setPerPage] = useState(10);

  useEffect(() => {
    axios.get('/product').then(({ data }) => setProducts(data));
  }, []);

  const onDelete = (productId: number) => {
    if (window.confirm('Are you sure you want to delete this product?')){
        axios.delete(`/product/${productId}`).then(() => {
            setProducts(products.filter(product => product.id !== productId))
        })
    }
  }

  const onEdit = (productId: number) => {
    navigate(`/product/${productId}/edit`);
  }

  return (
    <Box>
        <Button variant="contained" onClick={() => navigate('/product/create')}>Create Product</Button>
        <TableContainer sx={{ mt: 2 }}>
        <Table>
            <TableRow>
            <TableCell>#</TableCell>
            <TableCell>Image</TableCell>
            <TableCell>Title</TableCell>
            <TableCell>Description</TableCell>
            <TableCell>Price</TableCell>
            <TableCell>Action</TableCell>
            </TableRow>
            {products.slice(page*perPage, (page+1)*perPage).map((product: IProduct) => (
            <TableRow key={product.id}>
                <TableCell>{product.id}</TableCell>
                <TableCell><img src={product.image} width="50" /></TableCell>
                <TableCell>{product.title}</TableCell>
                <TableCell>{product.description}</TableCell>
                <TableCell>${product.price}</TableCell>
                <TableCell>
                    <ButtonGroup>
                        <Button color="warning" size="small" variant="contained" onClick={() => onEdit(product.id)}>Edit</Button>
                        <Button color="error" size="small" variant="contained" onClick={() => onDelete(product.id)}>Delete</Button>
                    </ButtonGroup>
                </TableCell>
            </TableRow>
            ))}
            <TableFooter>
            <TablePagination
                component="div"
                count={products.length}
                rowsPerPageOptions={[10, 20, 30]}
                onRowsPerPageChange={(e) => setPerPage(Number(e.target.value))}
                page={page}
                rowsPerPage={perPage}
                onPageChange={(_, page) => setPage(page)}
            />
            </TableFooter>
        </Table>
        </TableContainer>
    </Box>
  );
}
