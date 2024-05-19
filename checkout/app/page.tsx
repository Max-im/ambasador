import 'bootstrap/dist/css/bootstrap.min.css';
import styles from "./page.module.css";
export default function Home() {
  return (
    
    <main className={styles.main}>
      <div className="container">
        <div className="py-5 text-center">
          <h2>Checkout form</h2>
          <p className="lead">Ambassador invites you to buy the next products</p>
        </div>

        <div className="row">
          <div className="col-md-4 order-md-2 mb-4">
            <h4 className="d-flex justify-content-between align-items-center mb-3">
              <span className="text-muted">Products</span>
            </h4>
            <ul className="list-group mb-3">
              <li className="list-group-item d-flex justify-content-between lh-condensed">
                <div>
                  <h6 className="my-0">Product name</h6>
                  <small className="text-muted">Brief description</small>
                </div>
                <span className="text-muted">$12</span>
              </li>
              <li className="list-group-item d-flex justify-content-between">
                <span>Total (USD)</span>
                <strong>$20</strong>
              </li>
            </ul>
          </div>
          <div className="col-md-8 order-md-1">
            <h4 className="mb-3">Address</h4>
            <form className="needs-validation">
              <div className="row">
                <div className="col-md-6 mb-3">
                  <label htmlFor="firstName">First name</label>
                  <input type="text" className="form-control" id="firstName" placeholder="First name" required/> 
                </div>
                <div className="col-md-6 mb-3">
                  <label htmlFor="lastName">Last name</label>
                  <input type="text" className="form-control" id="lastName" placeholder="Last name" required />
                </div>
              </div>

              <div className="mb-3">
                <label htmlFor="email">Email <span className="text-muted"></span></label>
                <input type="email" className="form-control" id="email" placeholder="you@example.com" required/>
              </div>

              <div className="mb-3">
                <label htmlFor="address">Address</label>
                <input type="text" className="form-control" id="address" placeholder="1234 Main St" required />
              </div>

              <div className="row">
                <div className="col-md-5 mb-3">
                  <label htmlFor="country">Country</label>
                  <input type="text" className="form-control" id="country" placeholder="Country" required />
                </div>
                <div className="col-md-4 mb-3">
                  <label htmlFor="city">City</label>
                  <input type="text" className="form-control" id="city" placeholder="City" required />
                </div>
                <div className="col-md-3 mb-3">
                  <label htmlFor="zip">Zip</label>
                  <input type="text" className="form-control" id="zip" placeholder="Zip" required /> 
                </div>
              </div>
              <hr className="mb-4" />

              <h4 className="mb-3">Payment</h4>
              <button className="btn btn-primary btn-lg btn-block" type="submit">Checkout</button>
            </form>
          </div>
        </div>
      </div>
  </main>);
}
