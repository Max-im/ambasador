import React, { useEffect, useState } from 'react';
import axios from 'axios';
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Table,
  TableCell,
  TableContainer,
  TableRow,
  Typography,
} from '@mui/material';
import ExpandMore from '@mui/icons-material/ExpandMore';

export interface IOrder {
  id: number;
  name: string;
  email: string;
  total: number;
  order_items: IOrderItem[];
}

export interface IOrderItem {
  id: number;
  product_title: string;
  price: number;
  quantity: number;
  admin_revenue: number;
  ambassador_revenue: number;
}

export default function Order() {
  const [orders, setOrders] = useState<IOrder[]>([]);

  useEffect(() => {
    axios.get('/orders').then(({ data }) => setOrders(data));
  }, []);

  return (
    <>
      {orders.map((order) => (
        <Accordion sx={{ mt: 2 }} key={order.id}>
          <AccordionSummary expandIcon={<ExpandMore />}>
            <Typography variant="h6" gutterBottom>
              {order.name}
            </Typography>
          </AccordionSummary>
          <AccordionDetails>
            <Typography gutterBottom>Email: {order.email}</Typography>
            <Typography gutterBottom>Total: ${order.total}</Typography>
            <TableContainer sx={{ mt: 2 }}>
              <Table>
                <TableRow>
                  <TableCell>#</TableCell>
                  <TableCell>Product</TableCell>
                  <TableCell>Price</TableCell>
                  <TableCell>Quantity</TableCell>
                  <TableCell>Admin Revenue</TableCell>
                  <TableCell>Ambassador Revenue</TableCell>
                </TableRow>
                {order.order_items.map((orderItem) => (
                  <TableRow key={orderItem.id}>
                    <TableCell>{orderItem.id}</TableCell>
                    <TableCell>{orderItem.product_title}</TableCell>
                    <TableCell>${orderItem.price}</TableCell>
                    <TableCell>{orderItem.quantity}</TableCell>
                    <TableCell>${orderItem.admin_revenue}</TableCell>
                    <TableCell>${orderItem.ambassador_revenue}</TableCell>
                  </TableRow>
                ))}
              </Table>
            </TableContainer>
          </AccordionDetails>
        </Accordion>
      ))}
    </>
  );
}
