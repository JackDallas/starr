package lidarr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetQualityDefinition returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinition() ([]*QualityDefinition, error) {
	var definition []*QualityDefinition

	err := l.GetInto("v1/qualitydefinition", nil, &definition)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualitydefinition): %w", err)
	}

	return definition, nil
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := l.GetInto("v1/qualityprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// AddQualityProfile updates a quality profile in place.
func (l *Lidarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output QualityProfile

	err = l.PostInto("v1/qualityProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (l *Lidarr) UpdateQualityProfile(profile *QualityProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = l.Put("v1/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	err := l.GetInto("v1/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// GetMetadataProfiles returns the metadata profiles.
func (l *Lidarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	var profiles []*MetadataProfile

	err := l.GetInto("v1/metadataprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(metadataprofile): %w", err)
	}

	return profiles, nil
}

// GetQueue returns the Lidarr Queue.
func (l *Lidarr) GetQueue(maxRecords int) (*Queue, error) {
	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params.Set("sortKey", "timeleft")
	params.Set("sortDir", "asc")
	params.Set("pageSize", strconv.Itoa(maxRecords))
	params.Set("includeUnknownArtistItems", "true")

	var queue Queue

	err := l.GetInto("v1/queue", params, &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetSystemStatus returns system status.
func (l *Lidarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := l.GetInto("v1/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetTags returns all the tags.
func (l *Lidarr) GetTags() ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := l.GetInto("v1/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (l *Lidarr) AddTag(label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = l.PostInto("v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

// UpdateTag updates the label for a tag.
func (l *Lidarr) UpdateTag(tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = l.PutInto("v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// GetArtist returns an artist or all artists.
func (l *Lidarr) GetArtist(mbID string) ([]*Artist, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("mbId", mbID)
	}

	var artist []*Artist

	err := l.GetInto("v1/artist", params, &artist)
	if err != nil {
		return artist, fmt.Errorf("api.Get(artist): %w", err)
	}

	return artist, nil
}

// GetArtistByID returns an artist from an ID.
func (l *Lidarr) GetArtistByID(artistID int64) (*Artist, error) {
	var artist Artist

	err := l.GetInto("v1/artist/"+strconv.FormatInt(artistID, starr.Base10), nil, &artist)
	if err != nil {
		return &artist, fmt.Errorf("api.Get(artist): %w", err)
	}

	return &artist, nil
}

// AddArtist adds a new artist to Lidarr, and probably does not yet work.
func (l *Lidarr) AddArtist(artist *Artist) (*Artist, error) {
	body, err := json.Marshal(artist)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Artist

	err = l.PostInto("v1/artist", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(artist): %w", err)
	}

	return &output, nil
}

// UpdateArtist updates an artist in place.
func (l *Lidarr) UpdateArtist(artist *Artist) (*Artist, error) {
	body, err := json.Marshal(artist)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Artist

	err = l.PutInto("v1/artist/"+strconv.FormatInt(artist.ID, starr.Base10), params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(artist): %w", err)
	}

	return &output, nil
}

// GetAlbum returns an album or all albums if mbID is "" (empty).
// mbID is the music brainz UUID for a "release-group".
func (l *Lidarr) GetAlbum(mbID string) ([]*Album, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("ForeignAlbumId", mbID)
	}

	var albums []*Album

	err := l.GetInto("v1/album", params, &albums)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return albums, nil
}

// GetAlbumByID returns an album by DB ID.
func (l *Lidarr) GetAlbumByID(albumID int64) (*Album, error) {
	var album Album

	err := l.GetInto("v1/album/"+strconv.FormatInt(albumID, starr.Base10), nil, &album)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return &album, nil
}

// UpdateAlbum updates an album in place; the output of this is currently unknown!!!!
func (l *Lidarr) UpdateAlbum(albumID int64, album *Album) (*Album, error) {
	put, err := json.Marshal(album)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Album

	err = l.PutInto("v1/album/"+strconv.FormatInt(albumID, starr.Base10), params, put, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(album): %w", err)
	}

	return &output, nil
}

// AddAlbum adds a new album to Lidarr, and probably does not yet work.
func (l *Lidarr) AddAlbum(album *AddAlbumInput) (*Album, error) {
	if album.Releases == nil {
		album.Releases = make([]*AddAlbumInputRelease, 0)
	}

	body, err := json.Marshal(album)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Album

	err = l.PostInto("v1/album", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(album): %w", err)
	}

	return &output, nil
}

// GetCommands returns all available Lidarr commands.
func (l *Lidarr) GetCommands() ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := l.GetInto("v1/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Lidarr.
func (l *Lidarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	if cmd == nil || cmd.Name == "" {
		return nil, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	var output CommandResponse

	if err := l.PostInto("v1/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// GetHistory returns the last few items from the history endpoint.
func (l *Lidarr) GetHistory(maxRecords int) (*History, error) {
	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params.Set("pageSize", strconv.Itoa(maxRecords))

	var history History

	err := l.GetInto("v1/history", params, &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
}
