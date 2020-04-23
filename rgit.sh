git filter-branch --env-filter 'if [ "$GIT_AUTHOR_EMAIL" = "marco@ctdevelopers.com" ]; then
     GIT_AUTHOR_EMAIL=marco.milon@gmail.com;
     GIT_AUTHOR_NAME="Marco Milon";
     GIT_COMMITTER_EMAIL=$GIT_AUTHOR_EMAIL;
     GIT_COMMITTER_NAME="$GIT_AUTHOR_NAME"; fi' -- --all