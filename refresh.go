package plex

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (p *Plex) RefreshLibraries() error {
	query := fmt.Sprintf("%s/library/sections/all/refresh", p.URL)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return err
	}

	if resp.Status == ErrorInvalidToken {
		return errors.New("invalid token")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New(ErrorNotAuthorized)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("there was an error in the request")
	}

	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)

	return nil
}

func (p *Plex) RefreshLibrarySection(sectionKey string) error {
	query := fmt.Sprintf("%s/library/sections/%s/refresh", p.URL, sectionKey)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return err
	}

	if resp.Status == ErrorInvalidToken {
		return errors.New("invalid token")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New(ErrorNotAuthorized)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("there was an error in the request")
	}

	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)

	return nil
}

func (p *Plex) RefreshLibraryPath(sectionKey string, path string) error {
	query := fmt.Sprintf("%s/library/sections/%s/refresh?path=%s", p.URL, sectionKey, url.QueryEscape(path))

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return err
	}

	if resp.Status == ErrorInvalidToken {
		return errors.New("invalid token")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New(ErrorNotAuthorized)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("there was an error in the request")
	}

	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)

	return nil
}
