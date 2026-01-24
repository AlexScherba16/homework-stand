-- +goose Up
-- +goose StatementBegin
insert into public.profile(user_id,name,email) values
(1,'Nick','nick@gmail.com'),(2,'Alice','alice@gmail.com'),(3,'Bob','bob@gmail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from public.profile where user_id = any(array[1,2,3]);
-- +goose StatementEnd
