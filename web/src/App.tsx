import useSWR, { mutate } from "swr";
import "./App.css";
import { FeedPost, Post } from "./Post";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";
import gohper from "./../public/gohper.svg";
import { CreatePostForm } from "./CreatePostForm";
import { LoginPage } from "./LoginPage";

const rawAPIURL = import.meta.env.VITE_API_URL || "http://localhost:8080";
const normalizedAPIURL = rawAPIURL.replace(/\/+$/, "");

export const API_URL = normalizedAPIURL.endsWith("/v1")
  ? normalizedAPIURL
  : `${normalizedAPIURL}/v1`;

export const fetcher = (at: string) => async (url: string) => {
  const response = await fetch(API_URL + url, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${at}`,
    },
  });

  if (!response.ok) {
    throw new Error(`Request failed with status ${response.status}`);
  }

  return response.json();
};

function App() {
  const [cookies, setCookie] = useCookies(["at"]);
  const at = cookies.at;
  const feedKey = at ? "/users/feed" : null;

  const redirect = useNavigate();

  const { data, error, isLoading } = useSWR<{ data: FeedPost[] }>(
    feedKey,
    at ? fetcher(at) : null,
  );

  if (!at) return <LoginPage />;

  if (error) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Unable to load your feed</h2>
          <p>Please refresh the page or sign in again.</p>
        </div>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Loading your feed...</h2>
          <p>Gathering the latest updates from your network.</p>
        </div>
      </div>
    );
  }

  const posts = data?.data ?? [];
  const totalComments = posts.reduce(
    (total, post) => total + post.comments_count,
    0,
  );

  const handleLogout = () => {
    setCookie("at", "");
    redirect("/");
    return;
  };

  const reFetchData = () => {
    mutate("/users/feed");
  };

  const handleClickPost = (id: number) => () => redirect(`/post/${id}`);

  return (
    <div className="page-shell">
      <main className="app-layout">
        <nav className="topbar">
          <div className="brand">
            <img src={gohper} className="logo" alt="GopherSocial logo" />
            <div>
              <p className="brand-kicker">Social Hub for Developers</p>
              <h1>GopherSocial</h1>
            </div>
          </div>

          <button className="btn btn-secondary" onClick={handleLogout}>
            Log out
          </button>
        </nav>

        <section className="hero-panel">
          <div className="hero-copy">
            <h2>Build in public with your gopher crew</h2>
            <p>
              Share quick updates, publish thoughtful posts, and stay in sync
              with everyone you follow.
            </p>
          </div>

          <div className="hero-metrics">
            <article className="metric-card">
              <span className="metric-value">{posts.length}</span>
              <span className="metric-label">Posts in feed</span>
            </article>
            <article className="metric-card">
              <span className="metric-value">{totalComments}</span>
              <span className="metric-label">Total comments</span>
            </article>
          </div>
        </section>

        <section className="section-card">
          <div className="section-heading">
            <h3>Create a post</h3>
            <p>Start a new conversation with your community.</p>
          </div>
          <CreatePostForm onFetchPosts={reFetchData} />
        </section>

        <section className="section-card">
          <div className="section-heading">
            <h3>Latest updates</h3>
            <p>Follow the newest posts and jump into discussions quickly.</p>
          </div>

          <div className="posts">
            {posts.map((post, index) => (
              <Post
                key={post.id}
                post={post}
                onClick={handleClickPost(post.id)}
                index={index}
              />
            ))}
          </div>

          {posts.length === 0 && (
            <p className="empty-state">
              No posts yet. Start following someone or publish your first post.
            </p>
          )}
        </section>
      </main>
    </div>
  );
}

export default App;
