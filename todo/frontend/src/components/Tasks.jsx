import React, { useState, useEffect } from "react";
import { useAuth } from "../contexts/AuthContext";
import "./Tasks.css";

function Tasks() {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const { user, logout } = useAuth();

  useEffect(() => {
    loadTasks();
  }, []);

  const loadTasks = async () => {
    try {
      setLoading(true);
      const userTasks = await window.go.main.App.GetAllTasks(user.id);
      setTasks(userTasks || []);
    } catch (err) {
      setError("Ошибка при загрузке задач");
      console.error("Load tasks error:", err);
    } finally {
      setLoading(false);
    }
  };

  const addTask = async (e) => {
    e.preventDefault();
    if (!newTask.trim()) return;

    try {
      setLoading(true);
      const taskId = await window.go.main.App.Create(user.id, newTask.trim());
      if (taskId) {
        setNewTask("");
        loadTasks(); // Перезагружаем список задач
      }
    } catch (err) {
      setError("Ошибка при создании задачи");
      console.error("Add task error:", err);
    } finally {
      setLoading(false);
    }
  };

  const toggleTask = async (taskId, isDone) => {
    try {
      if (isDone) {
        await window.go.main.App.MarkAsUndone(taskId, user.id);
      } else {
        await window.go.main.App.MarkAsDone(taskId, user.id);
      }
      loadTasks(); // Перезагружаем список задач
    } catch (err) {
      setError("Ошибка при обновлении задачи");
      console.error("Toggle task error:", err);
    }
  };

  const deleteTask = async (taskId) => {
    if (!window.confirm("Вы уверены, что хотите удалить эту задачу?")) {
      return;
    }

    try {
      await window.go.main.App.DeleteTaskByID(taskId, user.id);
      loadTasks(); // Перезагружаем список задач
    } catch (err) {
      setError("Ошибка при удалении задачи");
      console.error("Delete task error:", err);
    }
  };

  const handleLogout = () => {
    logout();
  };

  return (
    <div className="tasks-container">
      <header className="tasks-header">
        <h1>Мои задачи</h1>
        <div className="user-info">
          <span>Привет, {user.username}!</span>
          <button onClick={handleLogout} className="logout-button">
            Выйти
          </button>
        </div>
      </header>

      <form onSubmit={addTask} className="add-task-form">
        <input
          type="text"
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          placeholder="Добавить новую задачу..."
          disabled={loading}
          className="task-input"
        />
        <button
          type="submit"
          disabled={loading || !newTask.trim()}
          className="add-button"
        >
          {loading ? "Добавление..." : "Добавить"}
        </button>
      </form>

      {error && <div className="error-message">{error}</div>}

      <div className="tasks-list">
        {loading && tasks.length === 0 ? (
          <div className="loading">Загрузка задач...</div>
        ) : tasks.length === 0 ? (
          <div className="no-tasks">У вас пока нет задач</div>
        ) : (
          tasks.map((task) => (
            <div
              key={task.id}
              className={`task-item ${task.is_done ? "completed" : ""}`}
            >
              <div className="task-content">
                <input
                  type="checkbox"
                  checked={task.is_done}
                  onChange={() => toggleTask(task.id, task.is_done)}
                  className="task-checkbox"
                />
                <span className="task-description">{task.description}</span>
              </div>
              <button
                onClick={() => deleteTask(task.id)}
                className="delete-button"
                title="Удалить задачу"
              >
                ×
              </button>
            </div>
          ))
        )}
      </div>

      {tasks.length > 0 && (
        <div className="tasks-stats">
          Всего задач: {tasks.length} | Выполнено:{" "}
          {tasks.filter((task) => task.is_done).length} | Осталось:{" "}
          {tasks.filter((task) => !task.is_done).length}
        </div>
      )}
    </div>
  );
}

export default Tasks;


