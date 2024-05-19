import 'bootstrap/dist/css/bootstrap.min.css';
import styles from "../page.module.css";

export default function Success() {
  return (
    <main className={styles.main}>
      <div className="container">
        <div className="py-5 text-center">
          <h2>Success</h2>
          <p className="lead">Your payment was successful</p>
        </div>
      </div>
    </main>
  )
}
