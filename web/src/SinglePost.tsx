import useSWR from "swr";
import { FeedPost } from "./Post";
import { useNavigate, useParams } from "react-router-dom";
import { useCookies } from "react-cookie";
import { fetcher } from "./App";

export const SinglePost = () => {
  const { postID } = useParams();
  const [cookies] = useCookies(["at"]);
  const at = cookies.at;

  const redirect = useNavigate();

  const { data, error, isLoading } = useSWR<{ data: FeedPost }>(
    "/posts/" + postID,
    at ? fetcher(at) : null,
  );

  if (!at) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Please log in</h2>
          <p>You need to sign in to view this post.</p>
        </div>
      </div>
    );
  }

  if (!postID) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Post not found</h2>
          <p>The requested post id is missing.</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Unable to load post</h2>
          <p>Please try again in a moment.</p>
        </div>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Loading post...</h2>
          <p>Fetching details and comments.</p>
        </div>
      </div>
    );
  }

  const post = data?.data;
  if (!post) {
    return (
      <div className="page-shell">
        <div className="status-card">
          <h2>Post not found</h2>
          <p>This post may have been removed.</p>
        </div>
      </div>
    );
  }

  const postDate = new Date(post.created_at).toLocaleDateString(undefined, {
    day: "2-digit",
    month: "long",
    year: "numeric",
  });

  return (
    <div className="page-shell">
      <main className="single-layout">
        <article className="single-post-card">
          <p className="eyebrow">Post details</p>
          <h1>{post.title || "Untitled post"}</h1>
          <p className="single-post-content">{post.content}</p>
          <div className="single-post-meta">
            <span>{postDate}</span>
            <span>{post.comments_count} comments</span>
          </div>
        </article>

        <section className="section-card">
          <div className="section-heading">
            <h3>Comments</h3>
            <p>Join the conversation around this post.</p>
          </div>

          <div className="comments">
            {post.comments && post.comments.length > 0 ? (
              post.comments.map((comment) => (
                <article key={comment.id} className="comment-card">
                  <p className="comment-author">
                    {comment.user?.username || "anonymous"}
                  </p>
                  <p className="comment-content">{comment.content}</p>
                  <p className="comment-date">
                    {new Date(comment.created_at).toLocaleDateString()}
                  </p>
                </article>
              ))
            ) : (
              <p className="empty-state">
                No comments yet. Be the first to respond.
              </p>
            )}
          </div>
        </section>

        <button className="btn btn-secondary" onClick={() => redirect("/feed")}>
          Back to feed
        </button>
      </main>
    </div>
  );
};
