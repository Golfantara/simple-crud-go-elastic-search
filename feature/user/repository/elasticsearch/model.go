package elasticsearch

import (
	"bytes"
	"context"
	"elasticsearch/feature/user"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
)

type UserRepository struct {
	client *elasticsearch.Client
}

func NewUserRepository(client *elasticsearch.Client) user.RepositoryElasticsearch {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) Save(user user.User) error {
    userID := strconv.Itoa(user.ID)
    
    userData, err := json.Marshal(user)
    if err != nil {
        return fmt.Errorf("error marshaling user data: %w", err)
    }

    res, err := r.client.Index(
        "users",
        bytes.NewReader(userData),
        r.client.Index.WithDocumentID(userID),
        r.client.Index.WithContext(context.Background()),
        r.client.Index.WithRefresh("true"),
    )
    
    if err != nil {
        return fmt.Errorf("error indexing document: %w", err)
    }
    defer res.Body.Close()

    if res.IsError() {
        return fmt.Errorf("error indexing document: %s", res.String())
    }

    return nil
}

func (r *UserRepository) FindByID(id string) (user.User, error) {
    res, err := r.client.Get(
        "users",
        id,
        r.client.Get.WithContext(context.Background()),
    )
    
    if err != nil {
        return user.User{}, fmt.Errorf("error getting document: %w", err)
    }
    defer res.Body.Close()

    if res.IsError() {
        if res.StatusCode == 404 {
            return user.User{}, errors.New("user not found")
        }
        return user.User{}, fmt.Errorf("error getting document: %s", res.String())
    }

    var result struct {
        Source user.User `json:"_source"`
    }
    
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return user.User{}, fmt.Errorf("error parsing response: %w", err)
    }

    return result.Source, nil
}

func (r *UserRepository) SearchUsers(query string) ([]user.User, error) {
    var searchQuery map[string]interface{}
    
    if query == "" {
        searchQuery = map[string]interface{}{
            "query": map[string]interface{}{
                "match_all": map[string]interface{}{},
            },
        }
    } else {
        searchQuery = map[string]interface{}{
            "query": map[string]interface{}{
                "multi_match": map[string]interface{}{
                    "query":  query,
                    "fields": []string{"name", "email", "address"},
                },
            },
        }
    }

    var buf bytes.Buffer
    if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
        return nil, fmt.Errorf("error encoding query: %w", err)
    }

    res, err := r.client.Search(
        r.client.Search.WithContext(context.Background()),
        r.client.Search.WithIndex("users"),
        r.client.Search.WithBody(&buf),
        r.client.Search.WithTrackTotalHits(true),
    )
    if err != nil {
        return nil, fmt.Errorf("error searching documents: %w", err)
    }
    defer res.Body.Close()

    if res.IsError() {
        return nil, fmt.Errorf("error searching documents: %s", res.String())
    }

    var result struct {
        Hits struct {
            Hits []struct {
                Source user.User `json:"_source"`
            } `json:"hits"`
        } `json:"hits"`
    }

    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("error parsing response: %w", err)
    }

    users := make([]user.User, len(result.Hits.Hits))
    for i, hit := range result.Hits.Hits {
        users[i] = hit.Source
    }

    return users, nil
}

func (r *UserRepository) Delete(id string) error {
    res, err := r.client.Delete(
        "users",
        id,
        r.client.Delete.WithContext(context.Background()),
        r.client.Delete.WithRefresh("true"),
    )
    if err != nil {
        return fmt.Errorf("error deleting document: %w", err)
    }
    defer res.Body.Close()

    if res.IsError() {
        if res.StatusCode == 404 {
            return errors.New("user not found")
        }
        return fmt.Errorf("error deleting document: %s", res.String())
    }

    return nil
}

func (r *UserRepository) Update(id string, user user.User) error {
    updateDoc := map[string]interface{}{
        "doc": user,
    }

    data, err := json.Marshal(updateDoc)
    if err != nil {
        return fmt.Errorf("error marshaling update data: %w", err)
    }

    res, err := r.client.Update(
        "users",
        id,
        bytes.NewReader(data),
        r.client.Update.WithContext(context.Background()),
        r.client.Update.WithRefresh("true"),
    )
    if err != nil {
        return fmt.Errorf("error updating document: %w", err)
    }
    defer res.Body.Close()

    if res.IsError() {
        if res.StatusCode == 404 {
            return errors.New("user not found")
        }
        return fmt.Errorf("error updating document: %s", res.String())
    }

    return nil
}