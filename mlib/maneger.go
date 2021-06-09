package mlib

import (
	"errors"
)

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManeger struct {
	musics []MusicEntry
}

func NewMusicManeger() *MusicManeger {
	return &MusicManeger{make([]MusicEntry, 0)}
}

func (m *MusicManeger) Len() int {
	return len(m.musics)
}

func (m *MusicManeger) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= m.Len() {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManeger) Find(name string) *MusicEntry {
	if m.Len() == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManeger) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManeger) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		return nil
	}
	removeMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removeMusic;
}

func (m *MusicManeger) RemoveByName(name string) *MusicEntry {
	for i, music := range m.musics {
		if music.Name == name {
			removeMusic := &m.musics[i]
			m.musics = append(m.musics[:i], m.musics[i+1:]...)
			return removeMusic
		}
	}
	return nil;
}