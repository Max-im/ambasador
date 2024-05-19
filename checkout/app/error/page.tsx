import 'bootstrap/dist/css/bootstrap.min.css';
import styles from "../page.module.css";

export default function Error() {
  return (
    <main className={styles.main}>
      <div className="container">
        <div className="py-5 text-center">
          <h2>Error</h2>
          <p className="lead">Could not process your payment</p>
        </div>
      </div>
    </main>
  )
}
